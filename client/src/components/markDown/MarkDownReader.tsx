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
			const rawMarkup = await marked(markdown) // Преобразуем Markdown в HTML
			setHtml(DOMPurify.sanitize(rawMarkup)) // Очищаем HTML и обновляем состояние
		}

		getMarkdown()
	}, [markdown]) // Запускаем эффект при изменении markdown

	return (
		<div
			className={s.markdown} // Добавьте класс для стилей, если нужно
			dangerouslySetInnerHTML={{ __html: html }}
		/>
	)
}

export default MarkdownRenderer
