import React, {Component} from 'react';
import {
    Route,
    NavLink,
    HashRouter
} from 'react-router-dom';
import Home from './components/Home';
import About from './components/About';
import Contact from './components/Contact';

class App extends Component {
    render() {
        return (
            <HashRouter>
                <div>
                    <div className='title'>
                        <h1>Ben Goro Photography</h1>
                    </div>
                    
                    <ul className='navigation'>
                        <li><NavLink exact to='/'>Home</NavLink></li>
                        <li><NavLink to='/about'>About</NavLink></li>
                        <li><NavLink to='/contact'>Contact</NavLink></li>
                    </ul>

                    <div className='content'>
                        <Route exact path='/' component={Home}/>
                        <Route path='/about' component={About}/>
                        <Route path='/contact' component={Contact}/>
                    </div>
                </div>
            </HashRouter>
        );
    }
}

export default App;
