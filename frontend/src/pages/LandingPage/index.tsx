import * as React from 'react';
import * as ReactDOM from "react-dom"
import "./index.css";
import logo from "./gator.png";
import { TopNav } from "./TopNav/TopNav";
import { SearchSuggestions } from "./SearchSuggestions/SearchSuggestions";
export default function LandingPage() {
  
  

  return (
    <div className="index">
      <h1 className='input'> HAPPY HAP'S YELP </h1>
      <TopNav />
      <img src={logo} className={logo} alt="logo" />
       <input type="text" placeholder="Find..." className='input_box'/> 
       <button className='input_submit' type='submit'> Go </button>
       <SearchSuggestions />
    </div>
  );
};