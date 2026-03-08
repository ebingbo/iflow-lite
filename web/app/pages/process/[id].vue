<script setup lang="ts">
import type { Assignment } from '~/types/assignment'
import { type Node, type NodeType, nodeTypeConfigs } from '~/types/node'
import type { Transition } from '~/types/transition'
import type { RoleOption, UserOption } from '~/types/user'

type WorkflowNode = Node & { _draft?: boolean }
type WorkflowEdge = Transition & { _draft?: boolean }

const route = useRoute()
const { getProcess } = useProcessApi()
const { addNode, updateNode, deleteNode: deleteNodeApi } = useNodeApi()
const { addTransition, deleteTransition } = useTransitionApi()
const { addAssignment, deleteAssignment } = useAssignmentApi()
const { listUsers } = useUserApi()
const { listRoles } = useRoleApi()
const toast = useToast()
const id = computed(() => Number(route.params.id))
const NODE_WIDTH = 160
const NODE_SPACING_X = 240
const NODE_SPACING_Y = 160

if (!Number.isInteger(id.value) || id.value <= 0) {
  throw createError({
    statusCode: 400,
    statusMessage: '无效的流程ID'
  })
}

const {
  data: detailData,
  status,
  error,
  refresh
} = await useAsyncData(
  `process-detail-${id.value}`,
  () => getProcess(id.value),
  {
    lazy: true
  }
)

const process = computed(() => detailData.value?.process ?? null)
const loading = computed(() => ['idle', 'pending'].includes(status.value))
const workflowLoading = computed(() => loading.value && nodes.value.length === 0 && edges.value.length === 0)
const workflowLoaded = computed(() => !workflowLoading.value && !error.value)
const processDescription = computed(() => process.value?.description || '流程查看与配置')
const processCode = computed(() => process.value?.code || '未配置编码')
const processStatusLabel = computed(() => process.value?.status === 1 ? '正常' : '禁用')
const processStatusColor = computed(() => process.value?.status === 1 ? 'success' : 'neutral')
const nodeTypeOptions = computed(() => {
  return Object.entries(nodeTypeConfigs).map(([value, config]) => ({
    label: config.name,
    value
  }))
})
const assignModeOptions = [
  { label: '单人处理(single)', value: 'single' },
  { label: '候选认领(candidate)', value: 'candidate' }
] as const

useSeoMeta({
  title: computed(() => process.value?.name ? `流程详情 - ${process.value.name}` : '流程详情'),
  description: '流程详情页面'
})

// 状态
const showConfig = ref(true)
const scale = ref(1)
const offset = ref({ x: 0, y: 0 })
const savingNode = ref(false)
const savingAssignment = ref(false)
const creatingNode = ref(false)
const deletingNode = ref(false)
const draftNodeSeed = ref(-1)
const lastPickedNodeType = ref<NodeType | null>(null)
const principalLoading = ref(false)
const principalKeyword = ref('')
const principalSearchTimer = ref<ReturnType<typeof setTimeout> | null>(null)
const assignmentUserID = ref('')
const assignmentRoleID = ref('')

// 节点数据
const nodes = ref<WorkflowNode[]>([])
const edges = ref<WorkflowEdge[]>([])
const assignments = ref<Assignment[]>([])
const userOptions = ref<UserOption[]>([])
const roleOptions = ref<RoleOption[]>([])

// 计算属性
const currentNodes = computed(() => nodes.value)
const currentEdges = computed(() => edges.value)
const nodeById = computed(() => new Map(currentNodes.value.filter(node => !!node.id).map(node => [node.id!, node])))
const viewBounds = computed(() => {
  if (!canvasRef.value || currentNodes.value.length === 0) {
    return { x: 0, y: 0, width: 1000, height: 800 }
  }

  const rect = canvasRef.value.getBoundingClientRect()
  return {
    x: -offset.value.x / scale.value,
    y: -offset.value.y / scale.value,
    width: rect.width / scale.value,
    height: rect.height / scale.value
  }
})

// 交互状态
const selectedNode = ref<WorkflowNode | null>(null)
const draggingNode = ref<{ node: WorkflowNode, offsetX: number, offsetY: number } | null>(null)
const connectingFrom = ref<WorkflowNode | null>(null)
const connectingLine = ref<{ x1: number, y1: number, x2: number, y2: number } | null>(null)
const isPanning = ref(false)
const panStart = ref({ x: 0, y: 0 })
const canvasRef = ref<HTMLDivElement | null>(null)
const configTab = ref('node')
const configTabItems = computed(() => [
  {
    label: '节点配置',
    value: 'node'
  },
  {
    label: '分派配置',
    value: 'assignment',
    disabled: !selectedNode.value || !!selectedNode.value._draft
  }
])
const userSelectItems = computed(() => userOptions.value.map(item => ({
  label: `${item.name} (${item.email})`,
  value: String(item.id)
})))
const roleSelectItems = computed(() => roleOptions.value.map(item => ({
  label: `${item.name} (${item.code})`,
  value: String(item.id)
})))

const viewBox = computed(() => {
  return `${viewBounds.value.x} ${viewBounds.value.y} ${viewBounds.value.width} ${viewBounds.value.height}`
})

// 工具函数
const isSupportedNodeType = (type: string): type is NodeType => {
  return type === 'start' || type === 'end' || type === 'user_task' || type === 'join'
}

const normalizeNodeType = (type: string | undefined): NodeType => {
  return type && isSupportedNodeType(type) ? type : 'user_task'
}

const getNodeHeaderClass = (type: string | undefined) => {
  const normalizedType = normalizeNodeType(type)
  return nodeTypeConfigs[normalizedType].canvasFill
}

const getNodeRingClass = (type: string | undefined) => {
  const normalizedType = normalizeNodeType(type)
  const classes: Record<NodeType, string> = {
    start: 'stroke-success/50',
    end: 'stroke-error/50',
    user_task: 'stroke-primary/50',
    join: 'stroke-warning/50'
  }
  return classes[normalizedType]
}

const getNodeHeight = (node: WorkflowNode): number => {
  if (node.type !== 'user_task') return 70
  const nodeAssignments = getNodeAssignments(node)
  const hasUser = nodeAssignments.some(item => item.principal_type === 'user')
  const hasRole = nodeAssignments.some(item => item.principal_type === 'role')
  const lineCount = Number(hasUser) + Number(hasRole)
  return lineCount > 0 ? 70 + lineCount * 16 : 70
}

const getNodeAssignments = (node: WorkflowNode): Assignment[] => {
  if (!node.id) return []
  return assignments.value
    .filter(item => item.node_id === node.id)
    .sort((a, b) => a.priority - b.priority || a.id - b.id)
}

const getPrincipalLabel = (item: Assignment): string => {
  if (item.principal_type === 'role') {
    const role = roleOptions.value.find(roleItem => roleItem.id === item.principal_id)
    return role ? `角色:${role.name}` : `角色#${item.principal_id}`
  }
  const user = userOptions.value.find(userItem => userItem.id === item.principal_id)
  return user ? `用户:${user.name}` : `用户#${item.principal_id}`
}

const getNodeAssigneeText = (node: WorkflowNode): string => {
  const item = getNodeAssignments(node).find(item => item.principal_type === 'user')
  return item ? getPrincipalLabel(item).replace('用户:', '') : ''
}

const getNodeRoleText = (node: WorkflowNode): string => {
  const item = getNodeAssignments(node).find(item => item.principal_type === 'role')
  return item ? getPrincipalLabel(item).replace('角色:', '') : ''
}

const getNodeAssignmentTextY = (node: WorkflowNode, type: 'user' | 'role'): number => {
  const hasUser = !!getNodeAssigneeText(node)
  const hasRole = !!getNodeRoleText(node)
  const lineCount = Number(hasUser) + Number(hasRole)
  const top = 45
  const height = getNodeHeight(node) - 53
  const centerY = top + height / 2
  if (lineCount <= 1) return centerY
  if (type === 'user') return centerY - 8
  return centerY + 8
}

const getEdgeNodes = (edge: WorkflowEdge) => {
  if (!edge.from_node_id || !edge.to_node_id) return null
  const fromNode = nodeById.value.get(edge.from_node_id)
  const toNode = nodeById.value.get(edge.to_node_id)
  if (!fromNode || !toNode) return null
  return { fromNode, toNode }
}

const getEdgePath = (fromNode: WorkflowNode, toNode: WorkflowNode): string => {
  const x1 = fromNode.x + NODE_WIDTH / 2
  const y1 = fromNode.y + getNodeHeight(fromNode)
  const x2 = toNode.x + NODE_WIDTH / 2
  const y2 = toNode.y
  const midY = (y1 + y2) / 2

  return `M ${x1} ${y1} L ${x1} ${midY} L ${x2} ${midY} L ${x2} ${y2}`
}
const getEdgeMidpoint = (fromNode: WorkflowNode, toNode: WorkflowNode): { x: number, y: number } => {
  const x1 = fromNode.x + NODE_WIDTH / 2
  const y1 = fromNode.y + getNodeHeight(fromNode)
  const x2 = toNode.x + NODE_WIDTH / 2
  const y2 = toNode.y
  return { x: (x1 + x2) / 2, y: (y1 + y2) / 2 }
}

const renderedEdges = computed(() => {
  return currentEdges.value.map((edge, index) => {
    const result = getEdgeNodes(edge)
    if (!result) return null
    const { fromNode, toNode } = result

    return {
      key: `${edge.from_node_id}-${edge.to_node_id}-${index}`,
      edge,
      path: getEdgePath(fromNode, toNode),
      midpoint: getEdgeMidpoint(fromNode, toNode)
    }
  }).filter((item): item is NonNullable<typeof item> => !!item)
})

const buildAutoLayout = (rawNodes: WorkflowNode[], rawEdges: WorkflowEdge[]) => {
  const nodeIds = rawNodes
    .map(item => item.id)
    .filter((item): item is number => typeof item === 'number')
  const idSet = new Set(nodeIds)
  const indegree = new Map<number, number>()
  const adjacency = new Map<number, number[]>()

  for (const nodeId of nodeIds) {
    indegree.set(nodeId, 0)
    adjacency.set(nodeId, [])
  }

  for (const edge of rawEdges) {
    const fromId = edge.from_node_id
    const toId = edge.to_node_id
    if (!fromId || !toId || !idSet.has(fromId) || !idSet.has(toId)) continue
    adjacency.get(fromId)?.push(toId)
    indegree.set(toId, (indegree.get(toId) ?? 0) + 1)
  }

  const roots = nodeIds.filter(nodeId => (indegree.get(nodeId) ?? 0) === 0)
  const queue = roots.length > 0 ? [...roots] : [...nodeIds.slice(0, 1)]
  const levelMap = new Map<number, number>()
  const visited = new Set<number>()

  for (const root of queue) {
    levelMap.set(root, 0)
    visited.add(root)
  }

  while (queue.length > 0) {
    const current = queue.shift()!
    const currentLevel = levelMap.get(current) ?? 0
    const nextNodes = adjacency.get(current) ?? []

    for (const next of nextNodes) {
      const nextLevel = currentLevel + 1
      if (!levelMap.has(next) || (levelMap.get(next) ?? 0) < nextLevel) {
        levelMap.set(next, nextLevel)
      }
      if (!visited.has(next)) {
        queue.push(next)
        visited.add(next)
      }
    }
  }

  let fallbackLevel = Math.max(...Array.from(levelMap.values(), levelMapValue => levelMapValue), -1) + 1
  for (const nodeId of nodeIds) {
    if (!levelMap.has(nodeId)) {
      levelMap.set(nodeId, fallbackLevel++)
    }
  }

  const groupedByLevel = new Map<number, number[]>()
  for (const nodeId of nodeIds) {
    const level = levelMap.get(nodeId) ?? 0
    const bucket = groupedByLevel.get(level) ?? []
    bucket.push(nodeId)
    groupedByLevel.set(level, bucket)
  }

  const layout = new Map<number, { x: number, y: number }>()
  const levels = Array.from(groupedByLevel.keys()).sort((a, b) => a - b)
  for (const level of levels) {
    const bucket = (groupedByLevel.get(level) ?? []).sort((a, b) => a - b)
    bucket.forEach((nodeId, idx) => {
      layout.set(nodeId, {
        x: 80 + idx * NODE_SPACING_X,
        y: 80 + level * NODE_SPACING_Y
      })
    })
  }

  return layout
}

const normalizeNodes = (rawNodes: WorkflowNode[], rawEdges: WorkflowEdge[]): WorkflowNode[] => {
  const autoLayout = buildAutoLayout(rawNodes, rawEdges)

  return rawNodes.map((item, index) => {
    const normalizedType = normalizeNodeType(item.type)
    const autoPosition = item.id ? autoLayout.get(item.id) : null
    const hasExplicitPosition = typeof item.x === 'number'
      && typeof item.y === 'number'
      && !(item.x === 0 && item.y === 0)
    const x = hasExplicitPosition ? item.x : (autoPosition?.x ?? 80 + (index % 4) * NODE_SPACING_X)
    const y = hasExplicitPosition ? item.y : (autoPosition?.y ?? 80 + Math.floor(index / 4) * NODE_SPACING_Y)
    return {
      ...item,
      id: item.id ?? Date.now() + index,
      type: normalizedType,
      x,
      y,
      assign_mode: item.assign_mode || 'single'
    }
  })
}

const normalizeTransitions = (rawEdges: WorkflowEdge[], normalizedNodes: WorkflowNode[]) => {
  const validNodeIds = new Set(normalizedNodes.map(item => item.id).filter((item): item is number => typeof item === 'number'))

  return rawEdges.filter((edge) => {
    return !!edge.from_node_id
      && !!edge.to_node_id
      && edge.from_node_id !== edge.to_node_id
      && validNodeIds.has(edge.from_node_id)
      && validNodeIds.has(edge.to_node_id)
  })
}

watch(
  () => detailData.value,
  async (payload) => {
    if (!payload) return

    const normalizedNodes = normalizeNodes(payload.nodes ?? [], payload.transitions ?? [])
    nodes.value = normalizedNodes
    edges.value = normalizeTransitions(payload.transitions ?? [], normalizedNodes)
    assignments.value = payload.assignments ?? []
    selectedNode.value = null

    await nextTick()
    resetView()
  },
  { immediate: true }
)

watch(
  () => selectedNode.value,
  (node) => {
    if (!node || node._draft) {
      configTab.value = 'node'
    }
  }
)

// 事件处理
const handleNodeMouseDown = (e: MouseEvent, node: WorkflowNode) => {
  e.preventDefault()
  if (e.shiftKey) {
    connectingFrom.value = node
  } else {
    // 注意:现在使用的是 SVG 坐标系,不需要考虑 scale
    draggingNode.value = {
      node,
      offsetX: e.clientX,
      offsetY: e.clientY
    }
  }
  selectedNode.value = node
}
const handleNodeMouseUp = (e: MouseEvent, node: WorkflowNode) => {
  e.preventDefault() // ✅ 防止默认行为

  if (connectingFrom.value && connectingFrom.value.id !== node.id) {
    const edgeList = edges.value
    const exists = edgeList.some(edge => edge.from_node_id === connectingFrom.value!.id && edge.to_node_id === node.id)

    if (!exists) {
      if (connectingFrom.value._draft || node._draft) {
        edges.value.push({
          from_node_id: connectingFrom.value.id,
          to_node_id: node.id,
          _draft: true
        })
      } else {
        addTransition({
          process_id: id.value,
          from_node_id: connectingFrom.value.id,
          to_node_id: node.id
        }).then((created) => {
          edges.value.push(created)
        }).catch((err: unknown) => {
          toast.add({
            title: '连线失败',
            description: err instanceof Error ? err.message : '创建连线失败',
            color: 'error'
          })
        })
      }
    }
  }

  // ✅ 清除所有状态
  draggingNode.value = null
  connectingFrom.value = null
  connectingLine.value = null
}
const handleMouseMove = (e: MouseEvent) => {
  if (draggingNode.value && canvasRef.value) {
    const rect = canvasRef.value.getBoundingClientRect()

    // 计算鼠标在 SVG 坐标系中的位置
    const svgX = (e.clientX - rect.left - offset.value.x) / scale.value
    const svgY = (e.clientY - rect.top - offset.value.y) / scale.value

    const nodeList = nodes.value
    const index = nodeList.findIndex(n => n.id === draggingNode.value!.node.id)
    if (index !== -1) {
      const currentNode = nodeList[index]
      if (!currentNode) return

      // 保持节点中心在鼠标位置
      nodeList[index] = {
        ...currentNode,
        x: svgX - NODE_WIDTH / 2,
        y: svgY - getNodeHeight(currentNode) / 2
      }
    }
  } else if (connectingFrom.value && canvasRef.value) {
    const rect = canvasRef.value.getBoundingClientRect()
    const svgX = (e.clientX - rect.left - offset.value.x) / scale.value
    const svgY = (e.clientY - rect.top - offset.value.y) / scale.value

    connectingLine.value = {
      x1: connectingFrom.value.x + 80,
      y1: connectingFrom.value.y + getNodeHeight(connectingFrom.value),
      x2: svgX,
      y2: svgY
    }
  } else if (isPanning.value) {
    offset.value = {
      x: e.clientX - panStart.value.x,
      y: e.clientY - panStart.value.y
    }
  }
}
const handleMouseUp = () => {
  // ✅ 彻底清除所有拖拽状态
  draggingNode.value = null
  isPanning.value = false
  if (!connectingFrom.value) {
    connectingLine.value = null
  }
}
const handleCanvasMouseDown = (e: MouseEvent) => {
  const target = e.target as HTMLElement
  const tagName = target.tagName.toLowerCase()
  if (target === canvasRef.value || tagName === 'svg' || tagName === 'rect') {
    isPanning.value = true
    panStart.value = {
      x: e.clientX - offset.value.x,
      y: e.clientY - offset.value.y
    }
    selectedNode.value = null
  }
}
// 操作函数
const addNewNode = async (type: NodeType) => {
  lastPickedNodeType.value = type
  const newNode: WorkflowNode = {
    id: draftNodeSeed.value,
    x: (300 - offset.value.x) / scale.value,
    y: (200 - offset.value.y) / scale.value,
    name: nodeTypeConfigs[type].name,
    tag: '',
    code: '',
    process_id: id.value,
    process_code: process.value?.code || '',
    description: null,
    type,
    created_by: 0,
    created_by_name: '',
    updated_by: 0,
    updated_by_name: '',
    created_at: '',
    updated_at: '',
    assign_mode: 'single',
    _draft: true
  }
  draftNodeSeed.value -= 1
  nodes.value.push(newNode)
  selectedNode.value = newNode
}
const deleteNode = async () => {
  if (!selectedNode.value) return
  if (deletingNode.value) return
  deletingNode.value = true

  const deletingNodeId = selectedNode.value.id
  if (!deletingNodeId) {
    deletingNode.value = false
    return
  }

  if (selectedNode.value._draft) {
    nodes.value = nodes.value.filter(n => n.id !== deletingNodeId)
    edges.value = edges.value.filter(
      e => e.from_node_id !== deletingNodeId && e.to_node_id !== deletingNodeId
    )
    assignments.value = assignments.value.filter(item => item.node_id !== deletingNodeId)
    selectedNode.value = null
    deletingNode.value = false
    return
  }

  const relatedEdges = edges.value.filter(
    edge => edge.from_node_id === deletingNodeId || edge.to_node_id === deletingNodeId
  )
  const relatedAssignments = assignments.value.filter(item => item.node_id === deletingNodeId)

  try {
    await Promise.all([
      ...relatedEdges.filter(edge => edge.id).map(edge => deleteTransition(edge.id!)),
      ...relatedAssignments.filter(item => item.id).map(item => deleteAssignment(item.id))
    ])
    await deleteNodeApi(deletingNodeId)
  } catch (err: unknown) {
    toast.add({
      title: '删除节点失败',
      description: err instanceof Error ? err.message : '删除节点失败',
      color: 'error'
    })
    deletingNode.value = false
    return
  }

  nodes.value = nodes.value.filter(n => n.id !== selectedNode.value!.id)
  edges.value = edges.value.filter(
    e => e.from_node_id !== selectedNode.value!.id && e.to_node_id !== selectedNode.value!.id
  )
  assignments.value = assignments.value.filter(item => item.node_id !== selectedNode.value!.id)
  selectedNode.value = null
  deletingNode.value = false
}
const deleteEdge = async (from: number, to: number) => {
  const target = edges.value.find(edge => edge.from_node_id === from && edge.to_node_id === to)
  if (!target) return

  if (target.id) {
    try {
      await deleteTransition(target.id)
    } catch (err: unknown) {
      toast.add({
        title: '删除连线失败',
        description: err instanceof Error ? err.message : '删除连线失败',
        color: 'error'
      })
      return
    }
  }

  edges.value = edges.value.filter(e => !(e.from_node_id === from && e.to_node_id === to))
}
const updateNodeField = <K extends keyof WorkflowNode>(field: K, value: WorkflowNode[K]) => {
  if (!selectedNode.value) return
  const nodeList = nodes.value
  const index = nodeList.findIndex(n => n.id === selectedNode.value!.id)
  if (index !== -1) {
    const currentNode = nodeList[index]
    if (!currentNode) return

    nodeList[index] = {
      ...currentNode,
      [field]: value
    }
  }
}

const saveNodeConfig = async () => {
  if (!selectedNode.value) return
  if (savingNode.value) return

  savingNode.value = true
  let current = selectedNode.value
  const isDraftBeforeSave = !!current._draft

  try {
    if (current._draft) {
      const name = current.name?.trim()
      const code = current.code?.trim()
      const tag = current.tag?.trim()
      if (!name || !code || !tag || !current.type) {
        toast.add({
          title: '请完善节点信息',
          description: '草稿节点创建前，名称/编码/标签/类型均为必填',
          color: 'warning'
        })
        savingNode.value = false
        return
      }

      creatingNode.value = true
      const created = await addNode({
        process_id: id.value,
        process_code: process.value?.code,
        tag,
        name,
        code,
        type: current.type,
        assign_mode: current.assign_mode || 'single',
        x: Math.round(current.x),
        y: Math.round(current.y),
        description: current.description || ''
      })
      creatingNode.value = false

      const oldNodeID = current.id
      const persistedNode: WorkflowNode = {
        ...current,
        ...created,
        id: created.id,
        x: current.x,
        y: current.y,
        _draft: false
      }

      nodes.value = nodes.value.map(item => item.id === oldNodeID ? persistedNode : item)
      edges.value = edges.value.map(edge => ({
        ...edge,
        from_node_id: edge.from_node_id === oldNodeID ? created.id : edge.from_node_id,
        to_node_id: edge.to_node_id === oldNodeID ? created.id : edge.to_node_id
      }))
      current = persistedNode
      selectedNode.value = persistedNode
    } else {
      await updateNode({
        id: current.id,
        tag: current.tag || current.name,
        assign_mode: current.assign_mode || 'single',
        x: Math.round(current.x),
        y: Math.round(current.y),
        description: current.description || ''
      })
    }

    const draftEdges = edges.value.filter(edge => !edge.id)
    for (const edge of draftEdges) {
      if (!edge.from_node_id || !edge.to_node_id) continue
      const fromNode = nodes.value.find(item => item.id === edge.from_node_id)
      const toNode = nodes.value.find(item => item.id === edge.to_node_id)
      if (!fromNode || !toNode || fromNode._draft || toNode._draft) continue

      const existsPersisted = edges.value.some(item =>
        !!item.id
        && item.from_node_id === edge.from_node_id
        && item.to_node_id === edge.to_node_id
      )
      if (existsPersisted) {
        edges.value = edges.value.filter(item => item !== edge)
        continue
      }

      const createdEdge = await addTransition({
        process_id: id.value,
        from_node_id: edge.from_node_id,
        to_node_id: edge.to_node_id
      })
      edges.value = edges.value.map(item => item === edge ? createdEdge : item)
    }

    toast.add({
      title: isDraftBeforeSave ? '创建成功' : '保存成功',
      description: isDraftBeforeSave ? '草稿节点已创建' : '节点配置已保存',
      color: 'success'
    })
  } catch (err: unknown) {
    toast.add({
      title: '保存失败',
      description: err instanceof Error ? err.message : '保存节点配置失败',
      color: 'error'
    })
  } finally {
    creatingNode.value = false
    savingNode.value = false
  }
}

const saveAssignmentConfig = async () => {
  if (!selectedNode.value) return
  if (selectedNode.value._draft) {
    toast.add({
      title: '请先创建节点',
      description: '草稿节点创建后才能保存分派',
      color: 'warning'
    })
    return
  }
  if (savingAssignment.value) return

  savingAssignment.value = true
  const current = selectedNode.value

  try {
    const currentNodeAssignments = assignments.value.filter(item => item.node_id === current.id)
    const mode = current.assign_mode || 'single'
    const targets: Array<{ principal_type: 'user' | 'role', principal_id: number }> = []
    const userID = Number((assignmentUserID.value || '').trim())
    const roleID = Number((assignmentRoleID.value || '').trim())

    if (Number.isInteger(userID) && userID > 0) {
      targets.push({
        principal_type: 'user',
        principal_id: userID
      })
    }
    if (mode !== 'single' && Number.isInteger(roleID) && roleID > 0) {
      targets.push({
        principal_type: 'role',
        principal_id: roleID
      })
    }

    for (const item of currentNodeAssignments) {
      await deleteAssignment(item.id)
    }

    const createdAssignments: Assignment[] = []
    for (let index = 0; index < targets.length; index += 1) {
      const target = targets[index]
      if (!target) continue
      createdAssignments.push(await addAssignment({
        process_id: id.value,
        node_id: current.id,
        principal_type: target.principal_type,
        principal_id: target.principal_id,
        priority: index,
        strategy: mode === 'candidate' ? 'parallel' : 'sequential'
      }))
    }

    assignments.value = assignments.value
      .filter(item => item.node_id !== current.id)
      .concat(createdAssignments)

    toast.add({
      title: '保存成功',
      description: '分派配置已保存',
      color: 'success'
    })
  } catch (err: unknown) {
    toast.add({
      title: '保存失败',
      description: err instanceof Error ? err.message : '保存分派配置失败',
      color: 'error'
    })
  } finally {
    savingAssignment.value = false
  }
}

const zoomIn = () => {
  if (!canvasRef.value) return
  const rect = canvasRef.value.getBoundingClientRect()
  const centerX = rect.width / 2
  const centerY = rect.height / 2

  const oldScale = scale.value
  const newScale = Math.min(scale.value + 0.1, 3)

  offset.value = {
    x: centerX - (centerX - offset.value.x) * (newScale / oldScale),
    y: centerY - (centerY - offset.value.y) * (newScale / oldScale)
  }

  scale.value = newScale
}

const zoomOut = () => {
  if (!canvasRef.value) return
  const rect = canvasRef.value.getBoundingClientRect()
  const centerX = rect.width / 2
  const centerY = rect.height / 2

  const oldScale = scale.value
  const newScale = Math.max(scale.value - 0.1, 0.2)

  offset.value = {
    x: centerX - (centerX - offset.value.x) * (newScale / oldScale),
    y: centerY - (centerY - offset.value.y) * (newScale / oldScale)
  }

  scale.value = newScale
}

const calculateBounds = () => {
  const padding = 100
  let minX = Infinity, minY = Infinity
  let maxX = -Infinity, maxY = -Infinity

  currentNodes.value.forEach((node) => {
    minX = Math.min(minX, node.x)
    minY = Math.min(minY, node.y)
    maxX = Math.max(maxX, node.x + 160)
    maxY = Math.max(maxY, node.y + getNodeHeight(node))
  })

  return {
    minX: minX - padding,
    minY: minY - padding,
    maxX: maxX + padding,
    maxY: maxY + padding,
    width: maxX - minX + padding * 2,
    height: maxY - minY + padding * 2
  }
}

const resetView = () => {
  if (!canvasRef.value || currentNodes.value.length === 0) {
    scale.value = 1
    offset.value = { x: 0, y: 0 }
    return
  }

  const bounds = calculateBounds()
  const rect = canvasRef.value.getBoundingClientRect()

  const scaleX = (rect.width * 0.9) / bounds.width
  const scaleY = (rect.height * 0.9) / bounds.height
  const newScale = Math.min(scaleX, scaleY, 1)

  const centerX = (bounds.minX + bounds.maxX) / 2
  const centerY = (bounds.minY + bounds.maxY) / 2

  offset.value = {
    x: rect.width / 2 - centerX * newScale,
    y: rect.height / 2 - centerY * newScale
  }

  scale.value = newScale
}

// ✅ 修复后的缩放函数
const handleWheel = (e: WheelEvent) => {
  if (!canvasRef.value) return
  e.preventDefault()

  const rect = canvasRef.value.getBoundingClientRect()
  const mouseX = e.clientX - rect.left
  const mouseY = e.clientY - rect.top

  const oldScale = scale.value
  const delta = e.deltaY > 0 ? -0.05 : 0.05
  const newScale = Math.max(0.2, Math.min(3, scale.value + delta))

  // 以鼠标位置为中心缩放
  offset.value = {
    x: mouseX - (mouseX - offset.value.x) * (newScale / oldScale),
    y: mouseY - (mouseY - offset.value.y) * (newScale / oldScale)
  }

  scale.value = newScale
}
const exportData = () => {
  const data = {
    processId: id.value,
    nodes: currentNodes.value,
    edges: currentEdges.value
  }
  console.log('流程配置:', data)
  useToast().add({
    title: '导出成功',
    description: '流程配置已导出到控制台',
    color: 'success'
  })
}

const retryLoad = async () => {
  await refresh()
}

const loadPrincipalOptions = async (keyword = '') => {
  principalLoading.value = true
  try {
    const [users, roles] = await Promise.all([
      listUsers({ size: 200, keyword }),
      listRoles({ size: 200, keyword })
    ])
    userOptions.value = users
    roleOptions.value = roles
  } catch (err: unknown) {
    toast.add({
      title: '加载分派选项失败',
      description: err instanceof Error ? err.message : '用户或角色列表加载失败',
      color: 'warning'
    })
  } finally {
    principalLoading.value = false
  }
}

onMounted(() => {
  loadPrincipalOptions()
})

const syncAssignmentSlots = () => {
  if (!selectedNode.value || selectedNode.value._draft || selectedNode.value.type !== 'user_task') {
    assignmentUserID.value = ''
    assignmentRoleID.value = ''
    return
  }
  const nodeAssignments = assignments.value
    .filter(item => item.node_id === selectedNode.value!.id)
    .sort((a, b) => a.priority - b.priority || a.id - b.id)
  const userAssignment = nodeAssignments.find(item => item.principal_type === 'user')
  const roleAssignment = nodeAssignments.find(item => item.principal_type === 'role')
  assignmentUserID.value = userAssignment ? String(userAssignment.principal_id) : ''
  assignmentRoleID.value = roleAssignment ? String(roleAssignment.principal_id) : ''
}

const assignmentSignature = computed(() => assignments.value
  .map(item => `${item.id}:${item.node_id}:${item.principal_type}:${item.principal_id}`)
  .join('|'))

watch(
  () => [selectedNode.value?.id, assignmentSignature.value],
  () => {
    syncAssignmentSlots()
  },
  { immediate: true }
)

watch(
  () => selectedNode.value?.assign_mode,
  (mode) => {
    if (!selectedNode.value || selectedNode.value.type !== 'user_task') return
    const assignMode = mode || 'single'
    if (assignMode === 'single') {
      assignmentRoleID.value = ''
    }
  }
)

watch(principalKeyword, (keyword) => {
  if (principalSearchTimer.value) {
    clearTimeout(principalSearchTimer.value)
  }
  principalSearchTimer.value = setTimeout(() => {
    loadPrincipalOptions(keyword.trim())
  }, 300)
})
</script>

<template>
  <UContainer class="space-y-4 pb-6">
    <UPageHeader
      :description="processDescription"
    >
      <template #title>
        <div class="flex items-center gap-2">
          <span>{{ process?.name || '流程详情' }}</span>
          <UBadge
            :color="processStatusColor"
            variant="soft"
            size="md"
          >
            {{ processStatusLabel }}
          </UBadge>
        </div>
      </template>
      <template #description>
        <div class="space-y-1 text-sm">
          <UBadge
            color="neutral"
            variant="soft"
          >
            编码: {{ processCode }}
          </UBadge>
          <div class="text-muted">
            {{ processDescription }}
          </div>
        </div>
      </template>
    </UPageHeader>

    <UAlert
      v-if="error"
      color="error"
      variant="subtle"
      title="加载失败"
      :description="error.message || '流程信息或流程图数据获取失败'"
      class="mb-4"
    >
      <template #actions>
        <UButton
          size="xs"
          color="error"
          variant="soft"
          @click="retryLoad"
        >
          重试
        </UButton>
      </template>
    </UAlert>

    <div
      v-else
      class="space-y-4"
    >
      <UCard>
        <div class="flex flex-wrap items-center justify-end gap-2">
          <UButton
            icon="i-lucide-minus"
            color="neutral"
            variant="soft"
            @click="zoomOut"
          />
          <UButton
            icon="i-lucide-scan"
            color="neutral"
            variant="soft"
            @click="resetView"
          />
          <UButton
            icon="i-lucide-plus"
            color="neutral"
            variant="soft"
            @click="zoomIn"
          />
          <UButton
            icon="i-lucide-download"
            color="success"
            variant="soft"
            @click="exportData"
          >
            导出
          </UButton>
          <UButton
            :icon="showConfig ? 'i-lucide-panel-right-close' : 'i-lucide-panel-right-open'"
            color="primary"
            variant="soft"
            @click="showConfig = !showConfig"
          >
            {{ showConfig ? '隐藏配置' : '显示配置' }}
          </UButton>
        </div>
      </UCard>

      <div class="flex h-[calc(100vh-280px)] min-h-[560px] gap-4">
        <div
          ref="canvasRef"
          class="relative flex-1 overflow-hidden rounded-xl border border-default bg-(--ui-bg-muted)"
          @mousedown="handleCanvasMouseDown"
          @mousemove="handleMouseMove"
          @mouseup="handleMouseUp"
          @wheel.prevent="handleWheel"
        >
          <div class="pointer-events-none absolute inset-0 bg-[radial-gradient(circle_at_20%_20%,rgba(59,130,246,0.08),transparent_40%),radial-gradient(circle_at_80%_0%,rgba(16,185,129,0.08),transparent_35%),linear-gradient(to_bottom,rgba(255,255,255,0.5),transparent_40%)] dark:bg-[radial-gradient(circle_at_20%_20%,rgba(56,189,248,0.12),transparent_40%),radial-gradient(circle_at_80%_0%,rgba(52,211,153,0.12),transparent_35%),linear-gradient(to_bottom,rgba(15,23,42,0.3),transparent_45%)]" />

          <div
            v-if="workflowLoading"
            class="absolute inset-0 z-20 flex items-center justify-center bg-(--ui-bg)/80"
          >
            <div class="text-sm text-muted">
              正在加载流程图...
            </div>
          </div>

          <div
            v-else-if="workflowLoaded && currentNodes.length === 0"
            class="absolute inset-0 z-20 flex items-center justify-center"
          >
            <div class="text-sm text-muted">
              当前流程暂无节点数据
            </div>
          </div>

          <UCard class="absolute left-4 top-4 z-10 w-56 border-default/80 bg-(--ui-bg)/90 shadow-lg backdrop-blur-sm">
            <template #header>
              <div class="space-y-1">
                <div class="flex items-center gap-2 font-semibold">
                  <UIcon
                    name="i-lucide-square-plus"
                    class="size-4 text-primary"
                  />
                  <span>添加节点</span>
                </div>
                <p class="text-xs text-muted">
                  选择类型后在画布生成草稿节点
                </p>
              </div>
            </template>
            <div class="space-y-2">
              <UButton
                v-for="(config, type) in nodeTypeConfigs"
                :key="type"
                block
                :icon="config.icon"
                :color="config.uiColor"
                variant="soft"
                class="justify-start transition-all duration-150"
                :class="lastPickedNodeType === type ? 'ring-1 ring-primary/60' : ''"
                :disabled="creatingNode"
                :loading="creatingNode"
                @click="addNewNode(type as NodeType)"
              >
                {{ config.name }}
              </UButton>
            </div>
          </UCard>

          <UAlert
            v-if="connectingFrom"
            color="warning"
            variant="soft"
            class="absolute left-1/2 top-4 z-10 w-auto -translate-x-1/2"
            :description="`从 “${connectingFrom.name}” 连线中，点击目标节点...`"
          />

          <UAlert
            class="absolute bottom-4 right-4 z-10 w-72"
            color="info"
            variant="soft"
            title="操作提示"
            icon="i-lucide-lightbulb"
          >
            <template #description>
              <ul class="list-disc space-y-1 text-xs">
                <li>拖拽节点可调整位置</li>
                <li>Shift + 点击开始连线</li>
                <li>点击 × 删除连线</li>
                <li>拖动画布可平移</li>
                <li>点击节点可编辑配置</li>
              </ul>
            </template>
          </UAlert>

          <UBadge
            class="absolute right-4 top-4 z-10"
            color="neutral"
            variant="soft"
            size="lg"
          >
            {{ Math.round(scale * 100) }}%
          </UBadge>

          <svg
            class="h-full w-full"
            :viewBox="viewBox"
          >
            <defs>
              <marker
                id="arrowhead"
                markerWidth="10"
                markerHeight="10"
                refX="9"
                refY="3"
                orient="auto"
              >
                <polygon
                  points="0 0, 10 3, 0 6"
                  fill="#94a3b8"
                />
              </marker>
              <pattern
                id="grid"
                width="20"
                height="20"
                patternUnits="userSpaceOnUse"
              >
                <circle
                  cx="1"
                  cy="1"
                  r="1"
                  fill="currentColor"
                  class="text-info/25"
                />
              </pattern>
            </defs>

            <rect
              :x="viewBounds.x"
              :y="viewBounds.y"
              :width="viewBounds.width"
              :height="viewBounds.height"
              fill="url(#grid)"
            />

            <g
              v-for="renderEdge in renderedEdges"
              :key="renderEdge.key"
              class="group"
            >
              <path
                :d="renderEdge.path"
                stroke="currentColor"
                class="stroke-info/55 transition-colors duration-150 group-hover:stroke-primary"
                stroke-width="2"
                fill="none"
                marker-end="url(#arrowhead)"
              />
              <circle
                :cx="renderEdge.midpoint.x"
                :cy="renderEdge.midpoint.y"
                r="8"
                fill="currentColor"
                class="fill-(--ui-bg) stroke-info/60 cursor-pointer transition-colors duration-150 group-hover:stroke-primary hover:fill-error/15"
                stroke-width="2"
                @click="deleteEdge(renderEdge.edge.from_node_id!, renderEdge.edge.to_node_id!)"
              />
              <text
                :x="renderEdge.midpoint.x"
                :y="renderEdge.midpoint.y"
                text-anchor="middle"
                dominant-baseline="middle"
                alignment-baseline="middle"
                class="fill-error text-xs font-bold leading-none pointer-events-none select-none"
              >
                ×
              </text>
            </g>

            <line
              v-if="connectingLine"
              :x1="connectingLine.x1"
              :y1="connectingLine.y1"
              :x2="connectingLine.x2"
              :y2="connectingLine.y2"
              stroke="currentColor"
              class="stroke-primary"
              stroke-width="2"
              stroke-dasharray="5,5"
            />

            <g
              v-for="node in currentNodes"
              :key="node.id"
              :transform="`translate(${node.x}, ${node.y})`"
              class="cursor-move"
              @mousedown.stop="handleNodeMouseDown($event, node)"
              @mouseup.stop="handleNodeMouseUp($event, node)"
            >
              <title
                v-if="node.type === 'user_task' && getNodeAssignments(node).length > 0"
              >
                {{ getNodeAssignments(node).map(getPrincipalLabel).join('\n') }}
              </title>

              <rect
                x="0"
                y="0"
                width="160"
                :height="getNodeHeight(node)"
                rx="8"
                fill="currentColor"
                class="fill-(--ui-bg)"
                :stroke="selectedNode?.id === node.id ? 'currentColor' : 'currentColor'"
                :class="selectedNode?.id === node.id ? 'stroke-primary' : 'stroke-muted'"
                :stroke-width="selectedNode?.id === node.id ? '3' : '2'"
                style="filter: drop-shadow(0 2px 8px rgba(0,0,0,0.1))"
              />
              <rect
                x="-4"
                y="-4"
                width="168"
                :height="getNodeHeight(node) + 8"
                rx="10"
                fill="none"
                stroke="currentColor"
                :class="getNodeRingClass(node.type)"
                stroke-width="1.5"
                class="pointer-events-none"
              />

              <rect
                x="0"
                y="0"
                width="160"
                height="35"
                rx="8"
                fill="currentColor"
                :class="getNodeHeaderClass(node.type)"
              />
              <rect
                x="0"
                y="8"
                width="160"
                height="27"
                fill="currentColor"
                :class="getNodeHeaderClass(node.type)"
              />

              <text
                x="80"
                y="23"
                text-anchor="middle"
                class="fill-white text-sm font-semibold"
              >
                {{ node.name }}
              </text>

              <g v-if="node.type === 'user_task' && getNodeAssignments(node).length > 0">
                <rect
                  x="8"
                  y="45"
                  width="144"
                  :height="getNodeHeight(node) - 53"
                  rx="4"
                  class="fill-info/10"
                />
                <text
                  v-if="getNodeAssigneeText(node)"
                  x="14"
                  :y="getNodeAssignmentTextY(node, 'user')"
                  dominant-baseline="middle"
                  alignment-baseline="middle"
                  class="fill-info text-[11px]"
                >
                  执行人: {{ getNodeAssigneeText(node) }}
                </text>
                <text
                  v-if="getNodeRoleText(node)"
                  x="14"
                  :y="getNodeAssignmentTextY(node, 'role')"
                  dominant-baseline="middle"
                  alignment-baseline="middle"
                  class="fill-info text-[11px]"
                >
                  执行角色: {{ getNodeRoleText(node) }}
                </text>
              </g>

              <circle
                cx="80"
                cy="0"
                r="6"
                fill="currentColor"
                :class="getNodeHeaderClass(node.type)"
                stroke="currentColor"
                class="stroke-(--ui-bg)"
                stroke-width="2"
                style="cursor: crosshair"
              />
              <circle
                cx="80"
                :cy="getNodeHeight(node)"
                r="6"
                fill="currentColor"
                :class="getNodeHeaderClass(node.type)"
                stroke="currentColor"
                class="stroke-(--ui-bg)"
                stroke-width="2"
                style="cursor: crosshair"
              />
            </g>
          </svg>
        </div>

        <div
          v-if="showConfig"
          class="w-80 overflow-y-auto"
        >
          <UCard>
            <UTabs
              v-model="configTab"
              :items="configTabItems"
              class="w-full"
            />

            <div
              v-if="configTab === 'node'"
              class="mt-4 space-y-4"
            >
              <div
                v-if="!selectedNode"
                class="py-12 text-center"
              >
                <svg
                  class="mx-auto mb-3 h-12 w-12 text-info/45"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
                  />
                </svg>
                <p class="text-muted">
                  点击节点进行配置
                </p>
              </div>

              <div
                v-else
                class="space-y-4"
              >
                <div class="flex items-center gap-2">
                  <UBadge
                    :color="selectedNode._draft ? 'warning' : 'neutral'"
                    variant="soft"
                  >
                    {{ selectedNode._draft ? '草稿节点' : `ID: ${selectedNode.id}` }}
                  </UBadge>
                  <UBadge
                    :icon="nodeTypeConfigs[normalizeNodeType(selectedNode.type)].icon"
                    :color="nodeTypeConfigs[normalizeNodeType(selectedNode.type)].uiColor"
                    variant="soft"
                  >
                    {{ nodeTypeConfigs[normalizeNodeType(selectedNode.type)].name }}
                  </UBadge>
                </div>

                <UFormField label="节点名称">
                  <UInput
                    v-model="selectedNode.name"
                    class="w-full"
                    :disabled="!selectedNode._draft"
                    @input="updateNodeField('name', selectedNode.name)"
                  />
                </UFormField>

                <UFormField label="节点编码">
                  <UInput
                    v-model="selectedNode.code"
                    class="w-full"
                    :disabled="!selectedNode._draft"
                    @input="updateNodeField('code', selectedNode.code)"
                  />
                </UFormField>

                <UFormField label="节点类型">
                  <USelect
                    v-model="selectedNode.type"
                    class="w-full"
                    :items="nodeTypeOptions"
                    value-key="value"
                    label-key="label"
                    :disabled="!selectedNode._draft"
                    @change="updateNodeField('type', selectedNode.type)"
                  />
                </UFormField>

                <UFormField
                  v-if="selectedNode.type === 'user_task'"
                  label="分派模式"
                >
                  <USelect
                    v-model="selectedNode.assign_mode"
                    class="w-full"
                    :items="assignModeOptions"
                    value-key="value"
                    label-key="label"
                    @change="updateNodeField('assign_mode', selectedNode.assign_mode)"
                  />
                </UFormField>

                <UFormField label="节点标签">
                  <UInput
                    v-model="selectedNode.tag"
                    class="w-full"
                    @input="updateNodeField('tag', selectedNode.tag)"
                  />
                </UFormField>

                <UFormField label="节点描述">
                  <UTextarea
                    v-model="selectedNode.description"
                    class="w-full"
                    :rows="3"
                    @input="updateNodeField('description', selectedNode.description)"
                  />
                </UFormField>

                <UButton
                  icon="i-lucide-save"
                  color="primary"
                  variant="solid"
                  block
                  :loading="savingNode"
                  :disabled="savingNode || savingAssignment || deletingNode"
                  @click="saveNodeConfig"
                >
                  {{ selectedNode._draft ? '创建节点' : '保存节点配置' }}
                </UButton>

                <UButton
                  icon="i-lucide-trash-2"
                  color="error"
                  variant="soft"
                  block
                  class="mt-2"
                  :loading="deletingNode"
                  :disabled="savingNode || savingAssignment || deletingNode"
                  @click="deleteNode"
                >
                  删除节点
                </UButton>
              </div>
            </div>

            <div
              v-else
              class="mt-4 space-y-4"
            >
              <div
                v-if="!selectedNode"
                class="py-8 text-sm text-muted"
              >
                请选择节点后配置分派
              </div>

              <div
                v-else-if="selectedNode._draft"
                class="py-8 text-sm text-muted"
              >
                请先在“节点配置”中创建节点，再保存分派
              </div>

              <div
                v-else-if="selectedNode.type !== 'user_task'"
                class="py-8 text-sm text-muted"
              >
                当前节点类型无需分派配置
              </div>

              <div
                v-else
                class="space-y-4"
              >
                <UInput
                  v-model="principalKeyword"
                  class="w-full"
                  placeholder="搜索用户姓名/邮箱/角色名称/编码"
                  icon="i-lucide-search"
                />

                <UFormField label="执行人">
                  <div class="flex items-center gap-2">
                    <USelect
                      v-model="assignmentUserID"
                      class="w-full"
                      :items="userSelectItems"
                      value-key="value"
                      label-key="label"
                      :loading="principalLoading"
                      placeholder="请选择执行人"
                    />
                    <UButton
                      icon="i-lucide-x"
                      color="neutral"
                      variant="soft"
                      :disabled="!assignmentUserID"
                      @click="assignmentUserID = ''"
                    />
                  </div>
                </UFormField>

                <UFormField label="执行角色">
                  <div class="flex items-center gap-2">
                    <USelect
                      v-model="assignmentRoleID"
                      class="w-full"
                      :items="roleSelectItems"
                      value-key="value"
                      label-key="label"
                      :loading="principalLoading"
                      placeholder="请选择执行角色"
                      :disabled="(selectedNode.assign_mode || 'single') === 'single'"
                    />
                    <UButton
                      icon="i-lucide-x"
                      color="neutral"
                      variant="soft"
                      :disabled="!assignmentRoleID || (selectedNode.assign_mode || 'single') === 'single'"
                      @click="assignmentRoleID = ''"
                    />
                  </div>
                </UFormField>

                <UButton
                  icon="i-lucide-users"
                  color="primary"
                  variant="soft"
                  block
                  :loading="savingAssignment"
                  :disabled="selectedNode._draft || savingAssignment || savingNode || deletingNode"
                  @click="saveAssignmentConfig"
                >
                  保存分派配置
                </UButton>
              </div>
            </div>
          </UCard>
        </div>
      </div>
    </div>
  </UContainer>
</template>
