const path = require('path');
const webpack = require('webpack');

module.exports = (
	(env) => {
		console.log(env);

		let local = env.local && env.local.length > 0 && (env.local === 'true' || env.local === true) ? true : false;

		return {
			watch: local,
			entry: './src/index.js',
			output: {
				filename: 'main.js',
				path: path.resolve(__dirname, 'dist'),
			},
			module: {
				rules: [
					{
						test: /\.(js|jsx)$/,
						exclude: /node_modules/,
						use: ['babel-loader']
					}
				]
			},
			resolve: {
				extensions: ['*', '.js', '.jsx']
			},
		};
	}
);