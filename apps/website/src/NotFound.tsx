import {Link} from "react-router";

export default function NotFound() {
    return (
        <>
            Sorry! We could not find that page, would you like to <Link to={"/"}>go Home</Link> instead?
        </>
    );
}
