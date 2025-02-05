export interface IProblem {
	ID: number
	title: string
	CreatedAt: string
	DeletedAt: string
	UpdatedAt: string
	body: string
	difficulty: number
	code: string
}
export interface IProblemSub {
	ID: number
	title: string
	CreatedAt: string
	DeletedAt: string
	UpdatedAt: string
	difficulty: number
}