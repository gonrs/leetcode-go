import ErrorPage from '@pages/helps/ErrorPage'
import LayOut from '@pages/helps/LayOut'
import Home from '@pages/home/Home'
import Problem from '@pages/problem/Problem'
import { createBrowserRouter } from 'react-router'

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
		],
	},
])
