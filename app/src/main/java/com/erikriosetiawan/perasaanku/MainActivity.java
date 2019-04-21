package com.erikriosetiawan.perasaanku;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.ImageView;
import android.widget.TextView;

public class MainActivity extends AppCompatActivity {

    ImageView imagePerasaan;
    TextView textPerasaan;
    boolean showingFist = true;


    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        getSupportActionBar().setTitle("Perasaan Saya");
        getSupportActionBar().setSubtitle("Ini adalah perasaan saya");
    }

    public void buttonClick(View view) {

        imagePerasaan = findViewById(R.id.img_happy);
        textPerasaan = findViewById(R.id.text_happy);

        if (showingFist) {
            imagePerasaan.setImageResource(R.drawable.img_tertawa);
            textPerasaan.setText(R.string.tertawa);
            showingFist = false;
        } else {
            imagePerasaan.setImageResource(R.drawable.img_menangis);
            textPerasaan.setText(R.string.menangis);
            showingFist = true;
        }
    }
}
