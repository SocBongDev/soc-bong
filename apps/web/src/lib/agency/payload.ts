import type { Paginator, Sort } from '$lib/common'

export class FindAgencyQuery {
	constructor(
		public Paginator: Paginator,
		public Sort: Sort,
		public Search: string,
		public OrderBy: string
	) {}

	/**
	 * Format
	 */
	public Format(): void {
		this.Paginator.Format()
		this.Sort.Format()
		this.Search = this.Search.trim()
		this.OrderBy = this.OrderBy.toLowerCase().trim()
	}
}
