import React from "react";
import Stack from "./Stack";
import StyledText from "./StyledText";

export default function Portfolio(props) {
	const [portfolio, setPortfolio] = useState({});
	
	return (
		<Stack>
			<StyledText styling='heading'>{portfolio.name}</StyledText>
		</Stack>
	);
}