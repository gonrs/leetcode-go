import { FC } from 'react'
import s from './style.module.css'
import { IProblem } from '@type/problemTypes'
import MarkdownRenderer from '@components/markDown/MarkDownReader'
interface IProblemDescrioptionProps {
	problem: IProblem | undefined
}
const ProblemDescrioption: FC<IProblemDescrioptionProps> = ({ problem }) => {
	if (!problem) {
		return <div className={s.problemBody}>Loading....</div>
	}
	let dif: string = 'easy'
	if (problem?.difficulty == 1) {
		dif = 'medium'
	} else if (problem?.difficulty == 2) {
		dif = 'hard'
	}
	return (
		<div className={s.problemBody}>
			<div className={s.problemBodyTitle}>
				<h1 className={s[dif]}>{problem.ID}</h1>
				<h2>{problem.title}</h2>
				<span className={s[dif]}>{dif}</span>
			</div>
			<div className={s.problemBodyDesc}>
				<MarkdownRenderer markdown={problem.body} />
			</div>
		</div>
	)
}
export default ProblemDescrioption
