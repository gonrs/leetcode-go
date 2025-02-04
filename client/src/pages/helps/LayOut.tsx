import React, { FC } from 'react'
import s from './style.module.css'
import Header from '@modules/header/Header'
import { Outlet } from 'react-router'
const LayOut: FC = ({}) => {
	return (
		<div className={s.layOutContainer}>
			<Header />
			<div className={s.container}>
				<Outlet />
			</div>
		</div>
	)
}
export default LayOut