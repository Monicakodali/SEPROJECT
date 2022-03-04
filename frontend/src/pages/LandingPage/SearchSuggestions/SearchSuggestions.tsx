import * as React from 'react';
import * as ReactDOM from "react-dom";
import styles from "./SearchSuggestions.module.css";
export function SearchSuggestions() {
    return (
	<div className={styles.suggestions}>
            <span className={styles.suggestion}>Restaurants</span>
	</div>
 );
};
