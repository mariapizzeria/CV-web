
import { EducationApi } from './components/api.js';
import { Renderer } from './components/renderer.js';

class App {
    private educationApi: EducationApi;
    private renderer: Renderer;

    constructor() {
        this.educationApi = new EducationApi();
        this.renderer = new Renderer('education-container');
    }

    async init(): Promise<void> {
        if (document.getElementById('education-container')) {
            await this.loadEducationData();
        }
    }

    private async loadEducationData(): Promise<void> {
        try {
            this.renderer.showLoading();
            const educationData = await this.educationApi.getEducation();
            this.renderer.createEducationItem(educationData);
        } catch(error) {
            console.error('Failed to load education data:', error);
            this.renderer.showLoading();
        }
    }
}

document.addEventListener('DOMContentLoaded', () => {
    const app = new App();
    app.init().catch(console.error);
});