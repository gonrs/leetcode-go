import { Link } from "react-router"
import s from './style.module.css'
import { FC } from "react"
import { URLS } from "@enums/URLS"
const Header: FC = ({}) => {
	return (
		<header className={s.header}>
			<Link to={URLS.home}>
				<h2>App</h2>
			</Link>
			<div>
        HUYYYUYUUY
			</div>
		</header>
	)
}
export default Header