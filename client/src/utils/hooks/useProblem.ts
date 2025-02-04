import { instance } from '@api/axios.api'
import { ServerURLS } from '@enums/URLS'
import { IProblem, IProblemSub } from '@type/problemTypes'
export const useProblem = () => {
	async function getProblems(
		from: number,
		to: number
	): Promise<IProblemSub[] | null> {
		try {
			const res = await instance.get<IProblemSub[]>(
				ServerURLS.getProblems + `?from=${from}&to=${to}`
			)
			if (res.status === 200) {
				console.log(res.data)
				return res.data
			}
			return null
		} catch (err) {
			console.error('Error checking server status:', err)
			return null
		}
	}
	async function getProblem(id: number): Promise<IProblem | null> {
		try {
			const res = await instance.get<IProblem>(ServerURLS.getProblem + id)
			if (res.status === 200) {
				console.log(res.data)
				return res.data
			}
			return null
		} catch (err) {
			console.log(err)
			return null
		}
	}
	return { getProblem, getProblems }
}