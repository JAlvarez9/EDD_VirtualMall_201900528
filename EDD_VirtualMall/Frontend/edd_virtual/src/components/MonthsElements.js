import React from 'react'
import { Dropdown } from 'semantic-ui-react'

function MonthsElements(props) {
    return (
        <>
            <Dropdown text={props.year} pointing className='link item'>
                <Dropdown.Menu >
                    <Dropdown.Header>Months</Dropdown.Header>
                    {props.months.map((c, index) => (
                        <Dropdown.Item> {c} </Dropdown.Item>
                    ))}
                </Dropdown.Menu>
            </Dropdown>

        </>
    )
}

export default MonthsElements
