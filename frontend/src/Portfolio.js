import React, { useEffect, useState } from "react";
import contactEndpoint from "./contactEndpoint";
import Stack from "./Stack";
import StyledText from "./StyledText";

export default function Portfolio(props) {
	const [portfolio, setPortfolio] = useState({});
	
	async function loadPortfolio(){
		try {
			let data = await contactEndpoint('GET', 'portfolio', {id: props.match.params.id});
			setPortfolio(data.Portfolio);
		}catch(e){
			// TODO: use a page error user can see
			console.log("Could not load portfolio: "+e);
		}
	}

	useEffect(
		() => {
			loadPortfolio();
		},
		[]
	)

	return (
		<Stack>
			<StyledText styling='heading'>{portfolio.Name}</StyledText>
		</Stack>
	);
}