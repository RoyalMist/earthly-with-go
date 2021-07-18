import {Link, Route} from "solid-app-router";


const App = () => {
    return (
        <>
            <div class="flex flex-row justify-around h-10 bg-blue-800">
                <Link class="text-white text-2xl" href="/">
                    Home
                </Link>
                <Link class="text-white text-2xl" href="/about">
                    About
                </Link>
            </div>

            <Route/>
        </>
    );
};

export default App;
