import { FC, useEffect, useState } from 'react'
import s from './style.module.css'
import { useProblem } from '@hooks/useProblem'
import { IProblemSub } from '@type/problemTypes'
import SubProblem from '@components/subProblem/SubProblem'

const Home: FC = ({}) => {
	const { getProblems } = useProblem()
	const [problemCoord, setProblemCoord] = useState({ from: 0, to: 10 })
	const [problems, setProblems] = useState<IProblemSub[]>([])

	async function getStartProblems() {
		try {
			const data = await getProblems(problemCoord.from, problemCoord.to)
			if (data != null) {
				setProblems(data)
			}
		} catch (err) {
			console.log(err)
		}
	}

	useEffect(() => {
		getStartProblems()
	}, [])
	return (
		<div>
			<div className={s.problems}>
				{problems.map((val, index) => {
					return <SubProblem title={val.title} id={val.ID} key={index} />
				})}
			</div>
			{/* <div className={s.setProblemsCoord}> */}
			{/* </div> */}
		</div>
	)
}

export default Home
