import { FC, useState } from 'react'
import s from './style.module.css'
import TextEditor from '@components/texxtEditor/TextEditor'
import { Button } from '@ui'
import { IProblem } from '@type/problemTypes'
import { instance } from '@api/axios.api'
import { ServerURLS } from '@enums/URLS'
import { ISendTest } from '@type/resTypes'
interface IProblemCodeEditor {
	problem: IProblem | undefined
	getRes: (res: ISendTest, resType: number) => void
	getIsSending: (res: boolean) => void
}
const ProblemCodeEditor: FC<IProblemCodeEditor> = ({
	problem,
	getRes,
	getIsSending,
}) => {
	if (!problem) {
		return <div className={s.problemBody}>Loading....</div>
	}
	const [isSending, setIsSending] = useState(false)
	const [code, setCode] = useState(problem.code)
	const [currentLanguage, setCurrentLanguage] = useState(0)
	async function send(type: number) {
		setIsSending(true)
		getIsSending(true)
		try {
			const data = await instance.post<ISendTest>(ServerURLS.sendTest, {
				problem_id: problem?.ID,
				code: code,
				type: type,
			})
			getRes(data.data, type)
		} catch (err) {
			console.log(err)
		}
		getIsSending(false)
		setIsSending(false)
	}
	function changeLanguage() {
		if (currentLanguage === 0) {
			setCurrentLanguage(1)
		} else {
			setCurrentLanguage(0)
		}
	}
	return (
		<div className={s.problemCodeEditor}>
			<TextEditor
				onChange={(code: string) => setCode(code)}
				startCode={problem.code}
			/>
			<div className={s.problemSend}>
				<Button disabled={isSending} onClick={() => send(0)}>
					Send Test
				</Button>
				<p>
					Language: <Button onClick={changeLanguage}>python</Button>
				</p>
				<Button disabled={isSending} onClick={() => send(1)}>
					Submite
				</Button>
			</div>
		</div>
	)
}
export default ProblemCodeEditor
