export class CSSTransition {
	static defaultProps: {
		classNames: string;
	};
	constructor(...args: any[]);
	forceUpdate(callback: any): void;
	reflowAndAddClass(node: any, className: any): void;
	removeClasses(node: any, type: any): void;
	render(): any;
	setState(partialState: any, callback: any): void;
}
export namespace CSSTransition {
	namespace propTypes {
		function addEndListener(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace addEndListener {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function appear(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace appear {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function children(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		function classNames(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace classNames {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function enter(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace enter {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function exit(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace exit {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function mountOnEnter(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace mountOnEnter {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function onEnter(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace onEnter {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function onEntered(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace onEntered {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function onEntering(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace onEntering {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function onExit(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace onExit {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function onExited(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace onExited {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function onExiting(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace onExiting {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function timeout(props: any, ...args: any[]): any;
		function unmountOnExit(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace unmountOnExit {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
	}
}
export class ReplaceTransition {
	constructor(...args: any[]);
	forceUpdate(callback: any): void;
	handleLifecycle(handler: any, idx: any, originalArgs: any): void;
	render(): any;
	setState(partialState: any, callback: any): void;
}
export namespace ReplaceTransition {
	namespace propTypes {
		function children(props: any, propName: any): any;
	}
}
export class Transition {
	static ENTERED: number;
	static ENTERING: number;
	static EXITED: number;
	static EXITING: number;
	static UNMOUNTED: number;
	static contextType: {
		$$typeof: symbol;
		Consumer: {
			$$typeof: symbol;
			Consumer: any;
			Provider: any;
		};
		Provider: {
			$$typeof: symbol;
		};
	};
	static getDerivedStateFromProps(_ref: any, prevState: any): any;
	constructor(props: any, context: any);
	cancelNextCallback(): void;
	componentDidMount(): void;
	componentDidUpdate(prevProps: any): void;
	componentWillUnmount(): void;
	forceUpdate(callback: any): void;
	getTimeouts(): any;
	onTransitionEnd(node: any, timeout: any, handler: any): void;
	performEnter(node: any, mounting: any): void;
	performExit(node: any): void;
	render(): any;
	safeSetState(nextState: any, callback: any): void;
	setNextCallback(callback: any): any;
	setState(partialState: any, callback: any): void;
	updateStatus(mounting: any, nextStatus: any): void;
}
export namespace Transition {
	namespace defaultProps {
		const appear: boolean;
		const enter: boolean;
		const exit: boolean;
		const mountOnEnter: boolean;
		function onEnter(): void;
		function onEntered(): void;
		function onEntering(): void;
		function onExit(): void;
		function onExited(): void;
		function onExiting(): void;
		const unmountOnExit: boolean;
	}
	namespace propTypes {
		function addEndListener(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace addEndListener {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function appear(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace appear {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function children(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		function enter(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace enter {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function exit(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace exit {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function mountOnEnter(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace mountOnEnter {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function onEnter(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace onEnter {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function onEntered(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace onEntered {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function onEntering(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace onEntering {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function onExit(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace onExit {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function onExited(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace onExited {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function onExiting(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace onExiting {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function timeout(props: any, ...args: any[]): any;
		function unmountOnExit(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace unmountOnExit {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
	}
}
export class TransitionGroup {
	static getDerivedStateFromProps(nextProps: any, _ref: any): any;
	constructor(props: any, context: any);
	componentDidMount(): void;
	componentWillUnmount(): void;
	forceUpdate(callback: any): void;
	handleExited(child: any, node: any): void;
	render(): any;
	setState(partialState: any, callback: any): void;
}
export namespace TransitionGroup {
	namespace defaultProps {
		function childFactory(child: any): any;
		const component: string;
	}
	namespace propTypes {
		function appear(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace appear {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function childFactory(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace childFactory {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function children(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace children {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function component(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace component {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function enter(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace enter {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
		function exit(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		namespace exit {
			function isRequired(p0: any, p1: any, p2: any, p3: any, p4: any, p5: any): any;
		}
	}
}
