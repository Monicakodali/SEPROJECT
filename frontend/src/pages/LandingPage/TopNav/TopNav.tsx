import * as React from 'react';
import * as ReactDOM from "react-dom"
import styles from "./TopNav.module.css";
//import { createBrowserHistory as history} from 'history';
import history from "./history";
import { Button } from 'react-bootstrap';
import Login from "../../../pages/LoginPage/index";
import { loginp } from "../../../pages/LoginPage/loginp"
export function TopNav() {
    return (
        <div className={styles['top-nav']}>
            <div className={styles.left}>
                <span>Write a Review</span>
                <span>Events</span>
            </div>
            <div className={styles.right}>
                <form>
                   <Button variant="btn btn-success" onClick={()=> history.push('/login')}> login </Button>
                </form>
                
                <button className='button'>Sign up</button>
            </div>
        </div>
    );
};
