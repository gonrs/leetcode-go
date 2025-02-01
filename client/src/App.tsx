import { RouterProvider } from 'react-router'
import './assets/App.css'
import { MainRoute } from '@routes/MainRoute.tsx'

function App() {
	return (
		<>
			<RouterProvider router={MainRoute} />
		</>
	)
}

export default App
