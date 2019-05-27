interface SignModelState {
    isLoggedIn: boolean
}
interface DvaEffect {
    call?: any
    put?: any
    takeEvery?: any
    select? :any
}

interface ModelState {
    sign: SignModelState
}   

export { DvaEffect, ModelState, SignModelState }
