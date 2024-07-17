import React, { useContext, useState } from "react";
import './content-component.css'
import { outputContext } from '../App.js';

function OutputComponent () {
    const myOutput =useContext(outputContext)
    
    return (
        <div>
            <pre>{JSON.stringify(myOutput, null, 2)}</pre>
        </div>
    )
}

export default OutputComponent;