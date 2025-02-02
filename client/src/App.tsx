import { RouterProvider } from 'react-router'
import './assets/App.css'
import { MainRoute } from '@routes/MainRoute.tsx'
import { useEffect } from 'react'
import { instance } from '@api/axios.api'
import { ServerURLS } from '@enums/URLS'

function App() {
	async function checkServerStatus() {
		try {
			const status = await instance.get(ServerURLS.checksServer)
			console.log(status)
		} catch (error) {
			console.error('Error checking server status:', error)
		}
	}
	useEffect(() => {
		checkServerStatus()
	}, [])

	return (
		<>
			<RouterProvider router={MainRoute} />
		</>
	)
}

export default App
