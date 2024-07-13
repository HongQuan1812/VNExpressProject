import React, { useState } from "react";
import './content-component.css'


function WhatButton () {
    const [myWhat, setmyWhat] = useState("news")

    const HandleClick = (e) => {
        setmyWhat(e.target.value)
    }

    const getButtonStyle = (value) => ({
        backgroundColor: myWhat === value ? 'green' : 'white',
        color: myWhat === value ? 'white' : 'black',
    });


    return (
        <div className="WhatButton" value={myWhat} >
            <button value="news" onClick={HandleClick} style={getButtonStyle("news")}> news </button>
            <button value="podcast" onClick={HandleClick} style={getButtonStyle("podcast")}> podcast </button>
            
        </div>   
    );
}

function ContentComponent () {
    const [myContent, setmyContent] = useState({})
    
    const HandleChange = (e, content) => {
        setmyContent(previousState => {
            return { ...previousState, [content] : e.target.value}
          });
    }
    return (
        <div className="ContentComponent">
            <WhatButton/> <br/>
            <h4> Main Categories </h4> 
            <input 
                type = "text"
                value = {myContent.MainCats}
                onChange={(e) => {HandleChange(e,"MainCats")}}
            />
            
            <h4> Sub Categories </h4>
            <input 
                type = "text"
                value = {myContent.SubCats}
                onChange={(e) => {HandleChange(e, "SubCats")}}
            />

            <h4> Authors </h4>
            <input 
                type = "text"
                value = {myContent.Authors}
                onChange={(e) => {HandleChange(e, "Authors")}}
            />

            <h4> Limit </h4>
            <input 
                type = "text"
                value = {myContent.Limit}
                onChange={(e) => {HandleChange(e, "Limit")}}
            />

            
            

        </div>
    );
}

export default ContentComponent;