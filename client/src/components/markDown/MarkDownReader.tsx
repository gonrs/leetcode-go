import { FC, useEffect, useState } from 'react'
import s from './style.module.css'
import { marked } from 'marked'
import DOMPurify from 'dompurify'
interface IMarkdownRendererProps {
	markdown: string
}
const MarkdownRenderer: FC<IMarkdownRendererProps> = ({ markdown }) => {
	const [html, setHtml] = useState<string>('')
	useEffect(() => {
		const getMarkdown = async () => {
			const rawMarkup = await marked(markdown)
			setHtml(DOMPurify.sanitize(rawMarkup)) 
		}
		getMarkdown()
	}, [markdown])
	return (
		<div
			className={s.markdown}
			dangerouslySetInnerHTML={{ __html: html }}
		/>
	)
}

export default MarkdownRenderer
