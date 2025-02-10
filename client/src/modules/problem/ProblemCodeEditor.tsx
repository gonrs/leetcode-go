import { FC, useEffect, useState } from 'react'
import s from './style.module.css'
import TextEditor from '@components/texxtEditor/TextEditor'
import { Button } from '@ui'
import { ILanguage, IProblem } from '@type/problemTypes'
import { instance } from '@api/axios.api'
import { ServerURLS } from '@enums/URLS'
import { ISendTest } from '@type/resTypes'
interface IProblemCodeEditor {
	problem: IProblem | undefined
	getRes: (res: ISendTest, resType: number) => void
	getIsSending: (res: boolean) => void
	setSolution: (res: boolean) => void
}
const ProblemCodeEditor: FC<IProblemCodeEditor> = ({
	problem,
	getRes,
	getIsSending,
	setSolution,
}) => {
	// if (!problem) {
	// 	return <div className={s.problemBody}>Loading....</div>
	// }
	const [isSending, setIsSending] = useState(false)
	const [languages, setLanguages] = useState<ILanguage[]>()
	const [currentLanguage, setCurrentLanguage] = useState(0)
	const [code, setCode] = useState('// some code')
	const [isSolution, setIsSolution] = useState(false)
	async function fetchLanguages() {
		if (problem) {
			try {
				const data = await instance.get<ILanguage[]>(
					ServerURLS.getLanguages + problem?.ID
				)
				setLanguages(data.data)
				setCode(data.data[currentLanguage].start_code)
			} catch (err) {
				console.log(err)
			}
		}
	}

	async function send(type: number) {
		if (languages) {
			setIsSending(true)
			getIsSending(true)
			try {
				const data = await instance.post<ISendTest>(ServerURLS.sendTest, {
					problem_id: problem?.ID,
					code: code,
					type: type,
					language_id: languages[currentLanguage].ID,
				})
				getRes(data.data, type)
			} catch (err) {
				console.log(err)
			}
			getIsSending(false)
			setIsSending(false)
		}
	}
	function changeLanguage() {
		if (currentLanguage === (languages ? languages.length - 1 : 0)) {
			setCurrentLanguage(0)
		} else {
			setCurrentLanguage(currentLanguage + 1)
		}
	}
	useEffect(() => {
		fetchLanguages()
	}, [problem])

	return problem ? (
		<div className={s.problemCodeEditor}>
			<TextEditor
				onChange={(code: string) => setCode(code)}
				startCode={
					languages ? languages[currentLanguage].start_code : '// some code'
				}
			/>
			<div className={s.problemSend}>
				{languages ? (
					<>
						<Button disabled={isSending} onClick={() => send(0)}>
							Send Test
						</Button>
						<p>
							Language:{' '}
							<Button onClick={changeLanguage}>
								{languages[currentLanguage].language}
							</Button>
						</p>
						<Button
							onClick={() => {
								setIsSolution(!isSolution)
								setSolution(!isSolution)
							}}
						>
							{isSolution ? 'Close Solution' : 'View Solution'}
						</Button>
						<Button disabled={isSending} onClick={() => send(1)}>
							Submit
						</Button>
					</>
				) : (
					<div className={s.problemBody}>Loading...</div>
				)}
			</div>
			
		</div>
	) : (
		<div className={s.problemBody}>Loading....</div>
	)
}
export default ProblemCodeEditor
