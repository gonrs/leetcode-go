export interface IProblem {
	ID: number
	title: string
	CreatedAt: string
	DeletedAt: string
	UpdatedAt: string
	body: string
	difficulty: number
	solution: string
}
export interface IProblemSub {
	ID: number
	title: string
	CreatedAt: string
	DeletedAt: string
	UpdatedAt: string
	difficulty: number
}
export interface ILanguage {
	ID: number
	CreatedAt: string
	DeletedAt: string
	problem_id: number
	language: string
	start_code: string
	help_code: string
}
