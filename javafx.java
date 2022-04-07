import javafx.embed.swing.JFXPanel;
import javafx.scene.Scene;
import javafx.scene.control.Button;
import javafx.scene.layout.StackPane;
import javafx.scene.paint.Color;

import javax.swing.*;
import java.awt.*;

public class Test {
    public static void main(String[] args) {
        // initialize a Swing application
        JFrame frame = new JFrame("JavaFX with Swing");

        // create a Swing compatible JavaFX container
        JFXPanel fxPanel = new JFXPanel();
        // initialize JavaFX container
        StackPane root = new StackPane();
        Scene scene = new Scene(root, 600, 400, Color.ALICEBLUE);
     
        // configure Swing JFrame
        frame.add(fxPanel);
        frame.setSize(600, 400);
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.setVisible(true);

     
    }
}
