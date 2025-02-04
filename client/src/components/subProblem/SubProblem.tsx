import { FC } from 'react'
import s from './style.module.css'
import { Link } from 'react-router'
import { URLS } from '@enums/URLS'
interface ISubProblemProps {
	title: string
	id: number
	difficulty: number
}
const SubProblem: FC<ISubProblemProps> = ({ id, title, difficulty }) => {
	let dif: string = 'easy'
	if (difficulty == 1) {
		dif = 'medium'
	} else if (difficulty == 2) {
		dif = 'hard'
	}
	return (
		<Link to={URLS.problem + id} className={s.mainButton}>
			<p className={s.id}>{id}</p>
			<p className={s.title}>{title}</p>
			<p className={s[dif]}>{dif}</p>
		</Link>
	)
}
export default SubProblem
