import { FC, useState } from 'react'
import s from './style.module.css'
import TextEditor from '@components/texxtEditor/TextEditor'
import { Button } from '@ui'
import { IProblem } from '@type/problemTypes'

interface IProblemCodeEditor {
	problem: IProblem | undefined
}
const ProblemCodeEditor: FC<IProblemCodeEditor> = ({ problem }) => {
	if (!problem) {
		return <div className={s.problemBody}>Loading....</div>
	}
	const [code, setCode] = useState(problem.startCode)
	function sendTest() {
		console.log(code)
	}
	function sendSubmite() {
		console.log(code)
	}
	return (
		<div className={s.problemCodeEditor}>
			<TextEditor
				onChange={(code: string) => setCode(code)}
				startCode={problem.startCode}
			/>
			<div className={s.problemSend}>
				<Button onClick={sendTest}>Send Test</Button>
				<p>
					Language: <span>python</span>
				</p>
				<Button onClick={sendSubmite}>Submite</Button>
			</div>
		</div>
	)
}

export default ProblemCodeEditor
