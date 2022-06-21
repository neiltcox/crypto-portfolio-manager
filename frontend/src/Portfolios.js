import React, { useEffect, useState } from "react";
import PortfolioTile from "./PortfolioTile";
import StyledText from "./StyledText";
import Stack from "./Stack";
import InlineLayout from "./InlineLayout";
import contactEndpoint from './contactEndpoint';

export default function Portfolios(props){
	const [portfolios, setPortfolios] = useState([]);

	async function loadPortfolios(){
		try {
			let data = await contactEndpoint('GET', 'portfolios');
			setPortfolios(data.Portfolios);
		}catch(e){
			console.log("could not load portfolios: "+e);
		}
	}

	useEffect(
		() => {
			loadPortfolios();
		},
		[]
	);

	return (
		<Stack>
			<InlineLayout>
				{
					portfolios.map(
						(portfolio) => {
							return (
								<React.Fragment key={portfolio.id}>
									<PortfolioTile portfolio={portfolio} />
								</React.Fragment>
							);
						}
					)
				}
			</InlineLayout>
		</Stack>
	);
}