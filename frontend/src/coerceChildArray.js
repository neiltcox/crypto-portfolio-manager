export default function coerceChildArray(children){
	return (
		(children &&
			(Array.isArray(children) &&
				// Children is an array, which is the format we want
				children
			) || (
				// Children is not an array, coerce into an array
				[children]
			)
		) || (
			// Children is null/undefined, use an empty array
			[]
		)
	).filter(
		(child) => {
			return child != null;
		}
	);
}