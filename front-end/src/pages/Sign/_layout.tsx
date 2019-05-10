import React from 'react'
import { TransitionGroup, CSSTransition } from 'react-transition-group'
import styles from './style.scss'
import Link from 'umi/link'

export default function Sign(props: any) {
    return (
        <div>
            <div className={styles['sign-page-header']}>
                <Link to='/home'>Logo</Link>
            </div>
            <TransitionGroup exit={false}>
                <CSSTransition key={location.pathname} classNames="fade" timeout={300}>
                    <div style={{ background: '#fff', padding: 24 }}>{props.children}</div>
                </CSSTransition>
            </TransitionGroup>
        </div>
    )
}
