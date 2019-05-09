import React from 'react'
import Link from 'umi/link'

export default () => {
    return (
        <div>
            I am a customized 404 page
            go back to <Link to="/home">Home</Link> 
        </div>
    )
}
