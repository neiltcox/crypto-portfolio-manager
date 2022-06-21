import React from "react";
import PortfolioTile from "./PortfolioTile";
import StyledText from "./StyledText";
import Stack from "./Stack";
import InlineLayout from "./InlineLayout";


export default function Portfolios(props){
	let portfolios = [
		{
			id: 1,
			name: "Long Term Index",
			exchange: "coinbasepro",
			valuation: 10230.83,
			connected: true,
		},
		{
			id: 2,
			name: "Speculative",
			exchange: "kraken",
			valuation: 6739.20,
			connected: false,
		}
	];

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