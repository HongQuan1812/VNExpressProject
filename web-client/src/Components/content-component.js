import React, { useContext, useEffect, useState } from "react";
import './content-component.css'
import { contentContext } from '../App.js';

function WhatButton () {
    const {myContent, setmyContent} = useContext(contentContext)

    const HandleClick = (event) => {
        setmyContent(prev => {
            const updatedContent = {...prev}
            updatedContent["myWhat"] = event.target.value
            return updatedContent
        })
    }
    const getButtonStyle = (value) => ({
        backgroundColor: myContent["myWhat"] === value ? 'green' : 'white',
        color: myContent["myWhat"] === value ? 'white' : 'black',
    });


    return (
        <div className="WhatButton" value={myContent["myWhat"]} >
            <button value="news" onClick={HandleClick} style={getButtonStyle("news")}> news </button>
            <button value="podcast" onClick={HandleClick} style={getButtonStyle("podcast")}> podcast </button>
            
        </div>   
    );
}

function ContentComponent () {
    const {myContent, setmyContent} = useContext(contentContext)

    const [inputValue, setinputValue] = useState({})
    const HandleChange = (event, content_name) => {
        setinputValue(prev => {
            let updatedContent = prev[content_name]
            updatedContent = event.target.value
            return {...prev, [content_name]: updatedContent}
          });
    }

    // handleBlur: Update myContent state when the input loses focus (onBlur event).
    const HandleBlur = (content_name) => {
        setmyContent(prev => {
            let updatedContent
            if (content_name === "Limit") {
                updatedContent = prev[content_name]
                updatedContent = inputValue[content_name].trim()
            } else {
                updatedContent = [...prev[content_name]]
                if (inputValue[content_name]) {
                    updatedContent = inputValue[content_name].split(/,\s*/)
                } else {
                    updatedContent = []
                }
            }
            return {...prev, [content_name]: updatedContent}
          }); 
    };

    useEffect(() => {
        setmyContent(prev => {
            return { ...prev, ["MainCats"]:[] }
        })
        setinputValue(prev => {
            return { ...prev, ["MainCats"]:[] }
        })
    }, [myContent["myWhat"]])

    return (
        <div className="ContentComponent">
            <WhatButton/> <br/>
            <h4> Main Categories </h4> 
            <input 
                type = "text"
                value = {myContent["myWhat"] === "podcast" ? "kinh tế" : inputValue["MainCats"]}
                onChange={(e) => HandleChange(e, "MainCats")}
                onBlur = {() => HandleBlur("MainCats")}
                disabled = {myContent["myWhat"] === "podcast" ? true : false}
            />
            
            <h4> Sub Categories </h4>
            <input 
                type = "text"
                value = {inputValue["SubCats"]}
                onChange={(e) => HandleChange(e, "SubCats")}
                onBlur = {() => HandleBlur("SubCats")}
            />

            <h4> Authors </h4>
            <input 
                type = "text"
                value = {inputValue["Authors"]}
                onChange={(e) => HandleChange(e, "Authors")}
                onBlur = {() => HandleBlur("Authors")}
            />

            <h4> Limit </h4>
            <input 
                type = "text"
                value = {inputValue["Limit"]}
                onChange={(e) => {HandleChange(e, "Limit")}}
                onBlur = {() => HandleBlur("Limit")}
            />

            
        </div>
    );
}

export default ContentComponent;