import React, { useState, useContext, useEffect } from "react";
import './when-component.css'
import { whenContext } from '../App.js';

function WhenComponent () {
    return (
        <div className = "WhenComponent">
            <DateComponent/>
            <TimeComponent/>
        </div>
    )
}

function YesNoButton () {
    const {myWhen, setmyWhen} = useContext(whenContext)
    
    const HandleClick = (e) => {
        setmyWhen(prev => {
            let updatedWholeDay = prev["WholeDay"]
            updatedWholeDay = (e.target.value === 'true')
            return {...prev, ["WholeDay"]: updatedWholeDay }
        });
    }

    const getButtonStyle = (value) => ({
        backgroundColor: myWhen["WholeDay"] === value ? 'green' : 'white',
        color: myWhen["WholeDay"] === value ? 'white' : 'black',
    });


    return (
        <div className="YesNoButton" value={myWhen["WholeDay"]} >
            <button value={true} onClick={HandleClick} style={getButtonStyle(true)}> yes </button>
            <button value={false} onClick={HandleClick} style={getButtonStyle(false)}> no </button>
            
        </div>   
    );
}

function Comparisor (props) {
    const {myWhen, setmyWhen} = useContext(whenContext)

    const handleChange = (index, event) => {
        setmyWhen(prev => {
            const updatedDayComparisor = [...prev["DayComparisor"]]
            updatedDayComparisor[index] = event.target.value

            if (event.target.value === "any") {
                return {...prev, ["DayComparisor"]:updatedDayComparisor , ["ReleaseDay"]: []}
            }
            if (event.target.value !== "BETWEEN") {
                const updatedReleaseDay = [...prev["ReleaseDay"]]
                if (updatedReleaseDay[0]) {
                    updatedReleaseDay[1] = "any-any-any"
                }
                return {...prev, ["DayComparisor"]:updatedDayComparisor, ["ReleaseDay"]: updatedReleaseDay}
            }
            return {...prev, ["DayComparisor"]:updatedDayComparisor }
        })
    };

    useEffect(() => {
        setmyWhen(prev => {
            return {... prev, ["DayComparisor"]:  ["any", "any", "any"]}
        })
    }, [myWhen["WholeDay"]])

    return (
        <select value = {myWhen["DayComparisor"][Number(props.index)]} onChange={(e) => {handleChange(Number(props.index), e)}} >
            <option value="any">any</option>
            <option value=">"> &gt; </option>
            <option value=">="> &gt;= </option>
            <option value="<"> &lt; </option>
            <option value="<="> &lt;= </option>
            <option value="="> = </option>
            <option value="BETWEEN"> BETWEEN </option>
        </select>
    );
}

function DateComponent () {
    const {myWhen, setmyWhen} = useContext(whenContext)

    const handleDate = (index, event) => {
        setmyWhen (prev => {
            const updatedReleaseDay = [...prev["ReleaseDay"]]
            if (event.target.value) {
                updatedReleaseDay[index] = event.target.value;
            } else {
                delete updatedReleaseDay[index];
            }
            
            if (myWhen["DayComparisor"][0] !== "BETWEEN") {
                if (updatedReleaseDay[0]) {
                    updatedReleaseDay[1] = "any-any-any";
                } else {
                    updatedReleaseDay.splice(0,2);
                }   
            } 
            else {
                if (!updatedReleaseDay[0] && !updatedReleaseDay[1]) {
                    updatedReleaseDay.splice(0,2);
                }
            }
            return {...prev, ["ReleaseDay"]: updatedReleaseDay}
        })
    };
        
    const date_part = ["Day", "Month", "Year"]
    const maxLength = [2,2,4]

    const [temp_date, settemp_date] = useState({
        "Day" : ["any", "any"],
        "Month" : ["any", "any"],
        "Year" : ["any", "any"],
    })

    const handleEachComponent = (element, index, event) => {
        settemp_date(prev => {
            const updatedElement = [...prev[element]];
            if (event.target.value) {
                updatedElement[index] = event.target.value;
            } else {
                updatedElement[index] = "any"
            }
            return { ...prev, [element]: updatedElement};
        });
    }

    useEffect(() => {
        let date1, date2
        date1 = [temp_date["Day"][0], temp_date["Month"][0],temp_date["Year"][0]].join('-')
        date2 = [temp_date["Day"][1], temp_date["Month"][1],temp_date["Year"][1]].join('-')
        if (temp_date["Day"][0] === "any" && temp_date["Month"][0] === "any" && temp_date["Year"][0] === "any") {
            date1 = undefined
        }   
        if (date1) {
            setmyWhen(prev => {
                return {...prev, ["ReleaseDay"]: [date1, date2]}        
            })
        } else {
            setmyWhen(prev => {
                return {...prev, ["ReleaseDay"]: []}        
            })
        }
    }, [temp_date]); // Run this effect whenever temp_date changes

    useEffect(() => {
        setmyWhen(prev => {
            return {... prev, ["ReleaseDay"]: []}
        })
        settemp_date (prev => {
            const updatedTemp_Date = {...prev}
            updatedTemp_Date["Day"] = ["any", "any"]
            updatedTemp_Date["Month"] = ["any", "any"]
            updatedTemp_Date["Year"] = ["any", "any"]
            return updatedTemp_Date
        })
    }, [myWhen["WholeDay"]])

    useEffect(() => {
        const handleKeyDown = (event) => {
            if (event.key === 'Enter') {
                event.target.blur();
            }
        };

        window.addEventListener('keydown', handleKeyDown);

        return () => {
            window.removeEventListener('keydown', handleKeyDown);
        };
    }, []);

    return (
        <div className="DateComponent">    
            <div className = "board">
                <h3> All Components in once ? </h3>
                <YesNoButton />
            </div>
            {
                myWhen["WholeDay"] ? (
                    <div className = "yesAnswer">
                        <h5> Comparisor </h5>
                        <Comparisor index = "0" />
                        <h5> Date </h5>
                        <div className="Date">
                            
                            <input
                                type="text"
                                value = {myWhen["DayComparisor"][0] === "any" || !myWhen["ReleaseDay"][0] ? "" : myWhen["ReleaseDay"][0]}
                                onChange={(e) => {handleDate(0,e)}}
                                placeholder="yyyy-mm-dd"
                                maxLength={10}
                                disabled={myWhen["DayComparisor"][0] === "any" ? true : false}
                            />
                            {
                                myWhen["DayComparisor"][0] === "BETWEEN" ? (
                                    <>
                                        <input
                                            type="text"
                                            value = {myWhen["ReleaseDay"][1]}
                                            onChange={(e) => {handleDate(1,e)}}
                                            placeholder="yyyy-mm-dd"
                                            maxLength={10}
                                        />
                                    </>
                                ):(
                                    null
                                )
                            }
                        </div>
                    </div>
                ) : (
                    <>
                        <ul>
                            {
                                
                                date_part.map((element, index) => { 
                                    return(
                                        <li className="noAnswer" key = {index}> 
                                            <h5> Comparisor </h5>
                                            <Comparisor index = {index} />
                                            <h5> {element} </h5>
                                            <div className="Day-Month-Year">
                                                
                                                <input
                                                    type="text"
                                                    placeholder={element}
                                                    maxLength={maxLength[index]}
                                                    value = {
                                                        (myWhen["DayComparisor"][index] === "any") || (temp_date[element][0] === "any") ? 
                                                        "" : 
                                                        temp_date[element][0]
                                                    }
                                                    onChange={(e) => {handleEachComponent(element, 0, e)}}
                                                    disabled={myWhen["DayComparisor"][index] === "any" ? true : false}
                                                />
                                                {
                                                    myWhen["DayComparisor"][index] === "BETWEEN" ? (
                                                        <>
                                                            <input
                                                                type="text"
                                                                placeholder={element}
                                                                maxLength={maxLength[index]}
                                                                value = {temp_date[element][1] === "any" ? 
                                                                    "" : 
                                                                    temp_date[element][1]}
                                                                onChange={(e) => {handleEachComponent(element, 1, e)}}
                                                            />
                                                        </>
                                                    ):(
                                                        null
                                                    )
                                                }
                                            </div>
                                        </li>
                                    )
                                })   
                            }
                        </ul>
                    </>
                    
                )
            }

            
        </div>
        

    );
    
}

function TimeComparisor() {
    const {myWhen, setmyWhen} = useContext(whenContext)

    const handleChange = (event) => {
        setmyWhen(prev => {
            let updatedTimeComparisor = prev["TimeComparisor"];
            updatedTimeComparisor = event.target.value
            if (event.target.value !== "BETWEEN") {
                const updatedReleaseTime = [...prev["ReleaseTime"]]
                if (updatedReleaseTime[0]) {
                    updatedReleaseTime[1] = "any:any"
                }
                return {...prev, ["TimeComparisor"]:updatedTimeComparisor , ["ReleaseTime"]: updatedReleaseTime}
            }
            if (event.target.value === "any") {
                return {...prev, ["TimeComparisor"]:updatedTimeComparisor, ["ReleaseTime"]:[]}
            }
            return {...prev, ["TimeComparisor"]: updatedTimeComparisor}
        })

    };
    return (
        <select value = {myWhen["TimeComparisor"]} onChange={handleChange} >
            <option value="any">any</option>
            <option value=">"> &gt; </option>
            <option value=">="> &gt;= </option>
            <option value="<"> &lt; </option>
            <option value="<="> &lt;= </option>
            <option value="="> = </option>
            <option value="BETWEEN"> BETWEEN </option>
        </select>
    );
}

function TimeComponent() {
    const {myWhen, setmyWhen} = useContext(whenContext)
    const handleChange = (index, event) => {
        setmyWhen(prev => {
            const updatedReleaseTime = [...prev["ReleaseTime"]]
            if (event.target.value) {
                updatedReleaseTime[index] = event.target.value
            } else {
                delete updatedReleaseTime[index]
            }
            
            if (myWhen["TimeComparisor"] !== "BETWEEN") {
                if (updatedReleaseTime[0]) {
                    updatedReleaseTime[1] = "any:any";
                } else {
                    updatedReleaseTime.splice(0,2);
                }   
            } 
            else {
                if (!updatedReleaseTime[0] && !updatedReleaseTime[1]) {
                    updatedReleaseTime.splice(0,2);
                }
            }

            return {...prev, ["ReleaseTime"]:updatedReleaseTime}
        })
    };

    return (
        <div className="TimeComponent">     
            <h3> How about time ? </h3>
            <div className = "Answer">
                <h5> Comparisor </h5>
                <TimeComparisor/>     
                <h5> Time </h5>
                <div className = "Time">
                    <input
                        type="text"
                        value = {myWhen["ReleaseTime"][0]}
                        onChange={(e) => {handleChange(0,e)}}
                        placeholder="hh:mm"
                        maxLength={10}
                        disabled={myWhen["TimeComparisor"] === "any" ? true : false}
                    />
                    {
                        myWhen["TimeComparisor"] === "BETWEEN" ? (
                            <>
                                <input
                                    type="text"
                                    value = {myWhen["ReleaseTime"][1]}
                                    onChange={(e) => {handleChange(1,e)}}
                                    placeholder="hh:mm"
                                    maxLength={10}
                                />
                            </>
                        ):(
                            null
                        )
                    }
                </div>
            </div>
        </div>
    );
}


export default WhenComponent;