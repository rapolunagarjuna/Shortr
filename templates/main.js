import * as THREE from 'https://unpkg.com/three@0.148.0/build/three.module.js';

// Create scene
const scene = new THREE.Scene();

// Create camera
const camera = new THREE.OrthographicCamera(
    window.innerWidth / -2,
    window.innerWidth / 2,
    window.innerHeight / 2,
    window.innerHeight / -2,
    0.1,
    1000
);
camera.position.z = 10;

// Create renderer
const renderer = new THREE.WebGLRenderer();
renderer.setSize(window.innerWidth, window.innerHeight);
document.body.appendChild(renderer.domElement);

// Resize handler
window.addEventListener('resize', () => {
    camera.left = window.innerWidth / -2;
    camera.right = window.innerWidth / 2;
    camera.top = window.innerHeight / 2;
    camera.bottom = window.innerHeight / -2;
    camera.updateProjectionMatrix();
    renderer.setSize(window.innerWidth, window.innerHeight);
});

// Render loop
function animate() {
    requestAnimationFrame(animate);
    renderer.render(scene, camera);
}
animate();

// Create a geometry and a material, then combine them into a mesh
const geometry = new THREE.BoxGeometry(100, 100, 0);
const material = new THREE.MeshBasicMaterial({ color: 0x00ff00 });
const cube = new THREE.Mesh(geometry, material);
scene.add(cube);

// Add event listener for mouse interaction
document.addEventListener('mousedown', onDocumentMouseDown, false);

function onDocumentMouseDown(event) {
    event.preventDefault();

    const mouse = new THREE.Vector2(
        (event.clientX / window.innerWidth) * 2 - 1,
        -(event.clientY / window.innerHeight) * 2 + 1
    );

    const raycaster = new THREE.Raycaster();
    raycaster.setFromCamera(mouse, camera);

    const intersects = raycaster.intersectObjects(scene.children);
    if (intersects.length > 0) {
        const selectedObject = intersects[0].object;
        selectedObject.material.color.set(0xff0000); // Change color on click
    }
}

let selectedObject = null;
let offset = new THREE.Vector3();

document.addEventListener('mousedown', onMouseDown, false);
document.addEventListener('mousemove', onMouseMove, false);
document.addEventListener('mouseup', onMouseUp, false);

function onMouseDown(event) {
    event.preventDefault();
    const mouse = new THREE.Vector2(
        (event.clientX / window.innerWidth) * 2 - 1,
        -(event.clientY / window.innerHeight) * 2 + 1
    );
    const raycaster = new THREE.Raycaster();
    raycaster.setFromCamera(mouse, camera);
    const intersects = raycaster.intersectObjects(scene.children);
    if (intersects.length > 0) {
        selectedObject = intersects[0].object;
        const intersection = intersects[0];
        offset.copy(intersection.point).sub(selectedObject.position);
    }
}

function onMouseMove(event) {
    if (selectedObject) {
        const mouse = new THREE.Vector2(
            (event.clientX / window.innerWidth) * 2 - 1,
            -(event.clientY / window.innerHeight) * 2 + 1
        );
        const raycaster = new THREE.Raycaster();
        raycaster.setFromCamera(mouse, camera);
        const intersects = raycaster.intersectObject(new THREE.Plane(new THREE.Vector3(0, 0, 1), 0));
        if (intersects.length > 0) {
            selectedObject.position.copy(intersects[0].point.sub(offset));
        }
    }
}

function onMouseUp(event) {
    selectedObject = null;
}
