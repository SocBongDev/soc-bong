export class Paginator {
	private page: number
	private pageSize: number

	constructor(urlSearchParams: URLSearchParams) {
		this.page = Number(urlSearchParams.get('page') ?? '0')
		this.pageSize = Number(urlSearchParams.get('pageSize') ?? '20')
	}

	public Page(): number {
		return this.page
	}

	public PageSize(): number {
		return this.pageSize
	}

	public SetPage(page: number): void {
		this.page = page
	}

	public SetPageSize(pageSize: number): void {
		this.pageSize = pageSize
	}

	/**
	 * Offset
	 */
	public Offset(): number {
		return (this.page - 1) * this.pageSize
	}

	/**
	 * Format
	 */
	public Format() {
		if (this.page <= 0) {
			this.page = 1
		}

		if (this.pageSize < 0) {
			this.pageSize = 20
		}

		if (this.pageSize > 200) {
			this.pageSize = 200
		}
	}
}
