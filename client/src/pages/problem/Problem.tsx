import { FC, useEffect, useState } from 'react'
import s from './style.module.css'
import { useParams } from 'react-router'
import { useProblem } from '@hooks/useProblem'
import { IProblem } from '@type/problemTypes'
import ProblemDescrioption from '@modules/problem/ProblemDescrioption'
import ProblemCodeEditor from '@modules/problem/ProblemCodeEditor'
import { ISendTest } from '@type/resTypes'
const Problem: FC = ({}) => {
	const { id } = useParams()
	const { getProblem } = useProblem()
	const [problem, setProblem] = useState<IProblem>()
	const [result, setResult] = useState<ISendTest | null>(null)
	const [sendType, setSendType] = useState<number>(0)
	async function getP() {
		try {
			const data = await getProblem(Number(id))
			if (data != null) {
				setProblem(data)
			} else {
				console.log('ERROR')
			}
		} catch (err) {
			console.log(err)
		}
	}
	useEffect(() => {
		getP()
	}, [])
	return (
		<div className={s.problem}>
			<div className={s.problemCon}>
				<ProblemDescrioption problem={problem} />
				<ProblemCodeEditor
					problem={problem}
					getRes={(res: ISendTest, type: number) => {
						setResult(res)
						setSendType(type)
					}}
				/>
			</div>
			{result &&
				(result.success ? (
					sendType === 0 ? (
						<div className={s.resultTestSucces}>
							<h1>Tests passed successfully</h1>
						</div>
					) : (
						<div className={s.resultSucces}>
							<h1>Succses</h1>
						</div>
					)
				) : (
					<div className={s.result}>
						<h1 className={s.resP}>
							Error in {result.test_index + 1}/{result.tests_length} test:
						</h1>
						<p className={s.resP}>Input: {result.test_input}</p>
						<p className={s.resP}>Output: {result.test_output}</p>
						<p className={s.resP}>Your output: {result.output}</p>
						<p className={s.resP}>Error: {result.error}</p>
					</div>
				))}
		</div>
	)
}
export default Problem
