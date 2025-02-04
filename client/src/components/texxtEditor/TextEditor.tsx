import React, { FC, useState } from 'react'
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
	const lines = text.split('\n')
	return (
		<div className={s.editorContainer}>
			<div className={s.lineNumbers}>
				{lines.map((_, index) => (
					<div key={index} className={s.lineNumber}>
						{index + 1}
					</div>
				))}
			</div>
			<textarea className={s.textArea} value={text} onChange={handleChange} />
		</div>
	)
}
export default TextEditor