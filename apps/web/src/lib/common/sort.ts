const ASC = 'ASC' as const
const DESC = 'DESC' as const

export class Sort {
	private sort: string

	constructor(urlSearchParams: URLSearchParams) {
		this.sort = urlSearchParams.get('sort') || DESC
	}

	/**
	 * IsValid
	 */
	public IsValid(): boolean {
		switch (this.sort) {
			case ASC:
			case DESC:
				return true
			default:
				return false
		}
	}

	public Format(): void {
		if (!this.IsValid()) {
			this.sort = DESC
		}
	}
}
