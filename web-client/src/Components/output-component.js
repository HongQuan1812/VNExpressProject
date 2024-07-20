import React, { useContext} from "react";
import './output-component.css'
import { outputContext } from '../App.js';

function OutputComponent () {
    const myOutput =useContext(outputContext)
    
    return (
        <div className="Output">
            <pre>{JSON.stringify(myOutput, null, 2)}</pre>
        </div>
    )
}

export default OutputComponent;