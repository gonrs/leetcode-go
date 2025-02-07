import { FC } from 'react'
import s from './style.module.css'
import { Link } from 'react-router'
const Home: FC = ({}) => {
	return (
		<div className={s.home}>
			<Link to={'/problemset'}>Problems</Link>
		</div>
	)
}
export default Home
