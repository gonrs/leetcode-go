import { FC } from 'react'
import s from './style.module.css'
import { useNavigate } from 'react-router'
import { URLS } from '@enums/URLS'

interface ISubProblemProps {
	title: string
	id: number
}

const SubProblem: FC<ISubProblemProps> = ({ id, title }) => {
	const navigate = useNavigate()
	async function HandleClick() {
		return navigate(URLS.problem + id)
	}
	return (
		<button onClick={HandleClick} className={s.mainButton}>
			<p className={s.id}>{id}</p>
			<p className={s.title}>{title}</p>
		</button>
	)
}

export default SubProblem
