import { FaSquareCaretRight } from "react-icons/fa6";
import { FileInput, Modal, Tabs, TextInput } from 'flowbite-react';
import { useRef, useState } from 'react';
import { useNavigate } from "react-router-dom";

const Scoring = (props: { role: any }) => {
    const navigate = useNavigate()
    const [openModal, setOpenModal] = useState(false);
    const [act, setAct] = useState("");

    const logout = () => {
        localStorage.removeItem("login")
        navigate("/login")
        window.location.reload()
    }

    return (
        <div>
            <p className="flex flex-row justify-end text-blue-500 hover:cursor-pointer active:text-blue-700" onClick={logout}>logout</p>
            <div className=" w-3/4 h-screen pt-24 m-auto flex flex-col items-start">
                <div className="flex flex-row" id="Title">
                    <p className="text-4xl">{props.role === 5 ? "Your Score" : "View Your Score"}</p>
                    <p className="text-4xl font-semibold ml-3" style={{ color: "#3078b3" }}>0</p>
                    <sub className="text-xl ml-1 mt-1" >pts</sub>
                </div>
                <div className="overflow-x-auto w-full">
                    <Tabs aria-label="Default tabs" className="flex focus:border-0" style="default">
                        <Tabs.Item active title="All Activities">
                            <div className="w-full">
                                <div className="flex flex-row mt-3">
                                    <FaSquareCaretRight style={{ color: '#3078b3', height: "25px" }} /> <p className="ml-2">All Activities</p>
                                </div>
                                <div id="Work-his" className=" mt-4 border border-2 border-slate-300 relative overflow-x-auto shadow-lg sm:rounded-lg ">
                                    <table className="w-full text-sm text-left text-center">
                                        <thead className=" text-white uppercase" style={{ background: "#3078b3" }}>
                                            <tr>
                                                <th>History No.</th>
                                                <th>Activities</th>
                                                <th>Score</th>
                                                <th>Supervisor</th>
                                            </tr>
                                        </thead>
                                        <tbody>
                                            <tr>
                                                <td>1</td>
                                                <td>Bowling</td>
                                                <td>0</td>
                                                <td>Hestia</td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </div>
                            </div>
                        </Tabs.Item>
                        <Tabs.Item title="Scored">
                            <div className="w-full">
                                <div className="flex flex-row mt-3">
                                    <FaSquareCaretRight style={{ color: '#3078b3', height: "25px" }} /> <p className="ml-2">Scored Activities</p>
                                </div>
                                <div id="Work-his" className=" mt-4 border border-2 border-slate-300 relative overflow-x-auto shadow-lg sm:rounded-lg ">
                                    <table className="w-full text-sm text-left text-center">
                                        <thead className=" text-white uppercase" style={{ background: "#3078b3" }}>
                                            <tr>
                                                <th>History No.</th>
                                                <th>Activities</th>
                                                <th>Score</th>
                                                <th>Supervisor</th>
                                            </tr>
                                        </thead>
                                        <tbody>
                                            <tr>
                                                <td>1</td>
                                                <td>Bowling</td>
                                                <td>0</td>
                                                <td>Hestia</td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </div>
                            </div>
                        </Tabs.Item>
                        <Tabs.Item title="Unscored">
                            <div className="w-full">
                                <div className="flex flex-row mt-3">
                                    <FaSquareCaretRight style={{ color: '#3078b3', height: "25px" }} /> <p className="ml-2">Unscored Activities</p>
                                </div>
                                <div id="Work-un" className=" mt-4 border border-2 border-slate-300 relative overflow-x-auto shadow-lg sm:rounded-lg ">
                                    <table className="w-full text-sm text-left text-center">
                                        <thead className=" text-white uppercase" style={{ background: "#3078b3" }}>
                                            <tr>
                                                <th>History No.</th>
                                                <th>Activities</th>
                                                <th>Score</th>
                                                <th>Supervisor</th>
                                            </tr>
                                        </thead>
                                        <tbody>
                                            <tr>
                                                <td>1</td>
                                                <td>Bowling</td>
                                                <td>0</td>
                                                <td>Hestia</td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </div>
                            </div>
                        </Tabs.Item>
                    </Tabs>
                    <button className="p-2 rounded-md mt-5 btn-def" onClick={() => setOpenModal(true)}>+ Add Activity</button>
                    <Modal show={openModal} onClose={() => setOpenModal(false)}>
                        <Modal.Header>Add Activity</Modal.Header>
                        <Modal.Body>
                            <div className="space-y-6">
                                <TextInput value={act} placeholder="Put your activity here" onChange={(e) => {
                                    if (e.target.value.length < 100) setAct(e.target.value)
                                    else alert("More than 100 characters")
                                }}></TextInput>
                            </div>
                        </Modal.Body>
                        <Modal.Footer>
                            <button className="p-2 rounded-md mt-5 btn-def" onClick={() => setOpenModal(false)}>I accept</button>
                            <button className="p-2 rounded-md mt-5 btn-def" style={{ color: "gray", backgroundColor: "white" }} onClick={() => {
                                setAct("")
                                setOpenModal(false)
                            }
                            }>
                                Decline
                            </button>
                        </Modal.Footer>
                    </Modal>
                </div>
            </div>
        </div>
    )
}

export default Scoring;