import { readBody } from 'h3'
import type { ApiResponse } from '~/types/api'
import type { TaskCandidateListBody, TaskCandidateListData } from '~/types/task'
import { backendFetch, ensureApiSuccess } from '~~/server/utils/backend'

export default eventHandler(async (event) => {
  const body = await readBody<TaskCandidateListBody>(event)

  const response = await backendFetch<ApiResponse<TaskCandidateListData>>(event, '/api/task/candidate/list', {
    method: 'POST',
    body: {
      task_id: body?.task_id
    }
  })

  return ensureApiSuccess(response, 'Task candidate list failed', 502)
})
