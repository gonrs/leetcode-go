import { FC, useEffect, useState } from 'react'
import s from './style.module.css'
import { useParams } from 'react-router'
import { useProblem } from '@hooks/useProblem'
import { IProblem } from '@type/problemTypes'

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

	let dif: string = 'easy'
	if (problem?.difficulty == 1) {
		dif = 'medium'
	} else if (problem?.difficulty == 2) {
		dif = 'hard'
	}

	return (
		<div>
			<h1>{problem?.ID}</h1>
			<h2>{problem?.title}</h2>
			<p>{problem?.body}</p>
			<p>{dif}</p>
		</div>
	)
}

export default Problem
