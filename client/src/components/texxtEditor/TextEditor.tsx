import React, { FC, useEffect, useRef, useState } from 'react'
import s from './style.module.css'
interface ITextEditorProps {
	startCode: string
	onChange: (value: string) => void
}
const TextEditor: FC<ITextEditorProps> = ({ startCode, onChange }) => {
	const [text, setText] = useState<string>(startCode)
	const handleChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
		setText(event.target.value)
		onChange(event.target.value)
	}
	const textareaRef = useRef<HTMLTextAreaElement>(null)
	const lineNumbersRef = useRef<HTMLDivElement>(null)
	const handleScroll = () => {
		console.log(1)
		if (textareaRef.current && lineNumbersRef.current) {
			lineNumbersRef.current.scrollTop = textareaRef.current.scrollTop
			console.log(2)
		}
	}
	useEffect(() => {
		const textarea = textareaRef.current
		const lineNumbers = lineNumbersRef.current
		if (textarea && lineNumbers) {
			const handleScroll = () => {
				lineNumbers.scrollTop = textarea.scrollTop
			}
			textarea.addEventListener('scroll', handleScroll)
			return () => {
				textarea.removeEventListener('scroll', handleScroll)
			}
		}
	}, [])
	const lines = text.split('\n')
	return (
		<div className={s.editorContainer}>
			<div ref={lineNumbersRef} className={s.lineNumbers}>
				{lines.map((_, index) => (
					<div key={index} className={s.lineNumber}>
						{index + 1}
					</div>
				))}
			</div>
			<textarea
				ref={textareaRef}
				className={s.textArea}
				value={text}
				onChange={handleChange}
				onScroll={handleScroll}
			/>
		</div>
	)
}
export default TextEditor
