import React from 'react'
import { TransitionGroup, CSSTransition } from 'react-transition-group'

export default function Sign(props: any) {
    return (
        <div>
            <TransitionGroup exit={false}>
                <CSSTransition key={location.pathname} classNames="fade" timeout={300}>
                    <div style={{ background: '#fff', padding: 24}}>
                        {props.children}
                    </div>
                </CSSTransition>
            </TransitionGroup>
        </div>
    )
}
