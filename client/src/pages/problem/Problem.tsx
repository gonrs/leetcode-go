import { FC, useEffect, useState } from 'react'
import s from './style.module.css'
import { useParams } from 'react-router'
import { useProblem } from '@hooks/useProblem'
import { IProblem } from '@type/problemTypes'
import ProblemDescrioption from '@modules/problem/ProblemDescrioption'

const Problem: FC = ({}) => {
	const { id } = useParams()

	const { getProblem } = useProblem()
	const [problem, setProblem] = useState<IProblem>()

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
		<div className={s.problemCon}>
			<ProblemDescrioption problem={problem}/>
		</div>
	)
}

export default Problem
