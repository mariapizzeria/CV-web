import { Api } from './components/api.js';
import { Renderer } from './components/renderer.js';
class App {
    constructor() {
        this.educationApi = new Api();
        this.skillApi = new Api();
        this.experienceApi = new Api();
        this.socialApi = new Api();
        this.renderer = new Renderer();
    }
    async init() {
        if (document.getElementById('education-container')) {
            await this.loadEducationData();
        }
        if (document.getElementById('skills-container')) {
            await this.loadSkillsData();
        }
        if (document.getElementById('experience-container')) {
            await this.loadExperienceData();
        }
        if (document.getElementById('social-container')) {
            await this.loadSocialData();
        }
    }
    async loadSocialData() {
        try {
            this.renderer.showLoading('social');
            const socialData = await this.socialApi.getSocial();
            this.renderer.createSocialItem(socialData);
        }
        catch (error) {
            console.error('Failed to load social data:', error);
            this.renderer.showError('Не удалось загрузить данные об соц.сетях', 'social');
        }
    }
    async loadEducationData() {
        try {
            this.renderer.showLoading('education');
            const educationData = await this.educationApi.getEducation();
            this.renderer.createEducationItem(educationData);
        }
        catch (error) {
            console.error('Failed to load education data:', error);
            this.renderer.showError('Не удалось загрузить данные об образовании', 'education');
        }
    }
    async loadSkillsData() {
        try {
            this.renderer.showLoading('skills');
            const skillsData = await this.skillApi.getSkills();
            this.renderer.createSkillItem(skillsData);
        }
        catch (error) {
            console.error('Failed to load skills data:', error);
            this.renderer.showError('Не удалось загрузить данные о навыках', 'skills');
        }
    }
    async loadExperienceData() {
        try {
            this.renderer.showLoading('experience');
            const experienceData = await this.experienceApi.getExperience();
            this.renderer.createExperienceItem(experienceData);
        }
        catch (error) {
            console.error('Failed to load experience data:', error);
            this.renderer.showError('Не удалось загрузить данные о навыках', 'experience');
        }
    }
}
document.addEventListener('DOMContentLoaded', () => {
    const app = new App();
    app.init().catch(console.error);
});
