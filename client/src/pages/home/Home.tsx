import React, { FC } from 'react'
import s from './style.module.css'
import { Link } from 'react-router'
import { Button } from '@ui'

const Home: FC = ({}) => {
	return (
		<div className={s.home}>
			<Button size='large'>
				<Link to={'/problemset'}>Problems</Link>
			</Button>
		</div>
	)
}

export default Home
