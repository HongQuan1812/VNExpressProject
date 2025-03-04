import { useState, createContext, useEffect} from "react";
import ContentComponent from './Components/content-component.js';
import WhenComponent from './Components/when-component.js';
import OutputComponent from './Components/output-component.js'
import {selectData} from './grpcClient.js';

import './App.css'

export const contentContext = createContext();
export const whenContext = createContext();
export const outputContext = createContext();

function App() {
  const [myContent, setmyContent] = useState({
      "myWhat": "news",
      "MainCats" : [],
      "SubCats" : [],
      "Authors" : [],
      "Limit" : "5",
  })
  const [myWhen, setmyWhen] = useState({
      "WholeDay" : true,
      "DayComparisor" : ["any", "any", "any"],
      "ReleaseDay" : [],
      "TimeComparisor" : "any",
      "ReleaseTime" : [],
  })

  const [myOutput, setmyOutput] = useState({})

  const [isButtonPressed, setIsButtonPressed] = useState(false);
  const [DownOnWhat, setDownOnWhat] = useState();

  const handleMouseDown = (event) => {
      setIsButtonPressed(true);
      setDownOnWhat(event.target);
      event.target.style.backgroundColor = "green";
  };
      
  useEffect(() => {
    const handleGlobalMouseUp = () => {
      if (isButtonPressed) {
        setIsButtonPressed(false);
        DownOnWhat.style.backgroundColor = "";
      }
    };
    window.addEventListener('mouseup', handleGlobalMouseUp);

    return () => {
      window.removeEventListener('mouseup', handleGlobalMouseUp);
    };
  }, [isButtonPressed]);
 
  
  const HandleClick = (event) => {
      selectData(
        myContent["myWhat"], // Directly pass the string value
        myContent["MainCats"], 
        myContent["SubCats"], 
        myContent["Authors"], 
        myWhen["WholeDay"], 
        myWhen["DayComparisor"], 
        myWhen["ReleaseDay"], 
        myWhen["TimeComparisor"], 
        myWhen["ReleaseTime"], 
        myContent["Limit"]
      ).then(result => {
        console.log("Fetched data:", result);
        setmyOutput(result)

      }).catch(error => {
        console.error("Error fetching data:", error);
        // Handle error, show error message, etc.
      });
  }
  
  
  return (
      <div className="App">
          <pre> {JSON.stringify(myContent, null, 2)} </pre>
          <pre> {JSON.stringify(myWhen, null, 2)} </pre>
          <div className="TableControl">
            <div className="TableContent">
                <contentContext.Provider value={{myContent, setmyContent}}>
                    <ContentComponent/>
                </contentContext.Provider>

                <whenContext.Provider value={{myWhen, setmyWhen}}>
                    <WhenComponent/>
                </whenContext.Provider>
            </div>
            <button 
                className="ConfirmButton" 
                onMouseDown={handleMouseDown}
                onClick={HandleClick}
            > Confirm 
            </button>
          </div>

          <div className="Output">
                <outputContext.Provider value={{myOutput}}>
                    <OutputComponent/>
                </outputContext.Provider>
          </div>

      </div>
      
      
    
  );
}

export default App;
