import { FC, useEffect, useState } from 'react'
import s from './style.module.css'
import { useProblem } from '@hooks/useProblem'
import { IProblemSub } from '@type/problemTypes'
import SubProblem from '@components/subProblem/SubProblem'
import { Button } from '@ui'

const Problemset: FC = ({}) => {
	const { getProblems } = useProblem()
	const [problemsPageLen, setProblemsPageLen] = useState(7)
	const [problemCoord, setProblemCoord] = useState({
		from: 0,
		to: problemsPageLen,
	})
	const [problems, setProblems] = useState<IProblemSub[]>([])

	async function getStartProblems(from: number, to: number) {
		try {
			const data = await getProblems(from, to)
			if (data != null) {
				setProblems(data)
			}
		} catch (err) {
			console.log(err)
		}
	}
	function handleClick(isPlus: boolean) {
		let newFrom = isPlus
			? problemCoord.from + problemsPageLen
			: problemCoord.from - problemsPageLen
		let newTo = isPlus
			? problemCoord.to + problemsPageLen
			: problemCoord.to - problemsPageLen

		if (newFrom < 0) {
			newFrom = 0
			newTo = problemCoord.to
		}
		setProblemCoord({
			from: newFrom,
			to: newTo,
		})
		getStartProblems(newFrom, newTo)
	}
	useEffect(() => {
		getStartProblems(problemCoord.from, problemCoord.to)
	}, [])
	return (
		<div>
			<div className={s.problems}>
				{problems.map((val, index) => {
					return (
						<SubProblem
							title={val.title}
							id={val.ID}
							key={index}
							difficulty={val.difficulty}
						/>
					)
				})}
			</div>

			<div className={s.setProblemsCoord}>
				{problemCoord.from != 0 ? (
					<Button onClick={() => handleClick(false)} size='medium'>
						prev
					</Button>
				) : (
					<></>
				)}
				{problems.length == problemCoord.to - problemCoord.from ? (
					<Button onClick={() => handleClick(true)} size='medium'>
						next
					</Button>
				) : (
					<></>
				)}
			</div>
		</div>
	)
}

export default Problemset
