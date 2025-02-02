import ErrorPage from '@pages/helps/ErrorPage'
import LayOut from '@pages/helps/LayOut'
import Problem from '@pages/problem/Problem'
import { createBrowserRouter } from 'react-router'
import Problemset from '@pages/problemset/Problemset'
import Home from '@pages/home/Home'

export const MainRoute = createBrowserRouter([
	{
		path: '/',
		element: <LayOut />,
		errorElement: <ErrorPage />,
		children: [
			{
				index: true,
				element: <Home />,
			},
			{
				path: 'problem/:id',
				element: <Problem />,
			},
			{
				path: 'problemset',
				element: <Problemset />,
			},
		],
	},
])
